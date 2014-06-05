package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	CONN_TYPE = "tcp"
)

var configarray []string
var logtwitter *bool

func main() {
	hostcfg := flag.String("listen", "localhost:6667", "<host>:<port>")
	logtwitter = flag.Bool("debuglog", false, "Enable if you want to log to a file")
	flag.Parse()
	configbytes, err := ioutil.ReadFile("./twitterauth.cfg")
	if err != nil {
		log.Fatal("Could not read the config file. not going to bother")
	}

	configarray = strings.Split(strings.Replace(string(configbytes), "\r", "", -1), "\n")
	if len(configarray) != 2 && (len(configarray) != 3) {
		if len(configarray) == 3 && configarray[2] == "" {

		} else {
			log.Fatal("bad amount of data in config.")
		}
	}

	if configarray[0] == "API key" {
		log.Fatal("you need to fill out the config files...")
	}

	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, *hostcfg)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + *hostcfg)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		// Handle connections in a new goroutine.
		go handleIRCConn(conn)
	}
}

// Handles incoming requests.
func handleIRCConn(conn net.Conn) {
	var ConnectionStage int = 0
	var TwitterToken string
	var IRCUsername string
	hostname, e := os.Hostname()
	if e != nil {
		hostname = "Unknown"
	}

	c := oauth.NewConsumer(
		configarray[0],
		configarray[1],
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	logindata := oauth.AccessToken{}

	reader := bufio.NewReader(conn)
	for {
		lineb, _, err := reader.ReadLine()
		line := string(lineb)

		if err != nil {
			return
		}

		fmt.Println(line)

		if strings.HasPrefix(line, "PASS ") && ConnectionStage == 0 {
			TwitterToken = strings.Split(line, " ")[1]
			json.Unmarshal([]byte(TwitterToken), &logindata)
			fmt.Printf("Twitter token: %s \n", TwitterToken)
			ConnectionStage++

		}

		if strings.HasPrefix(line, "NICK ") && ConnectionStage == 1 {
			fmt.Println(line)
			IRCUsername = strings.Split(line, " ")[1]
			conn.Write(GenerateIRCMessageBin(RplWelcome, IRCUsername, ":Welcome to TwiRC"))

			conn.Write(GenerateIRCMessageBin(RplYourHost, IRCUsername, fmt.Sprintf(":Host is: %s", hostname)))
			conn.Write(GenerateIRCMessageBin(RplCreated, IRCUsername, ":This server was first made on 31/06/2014"))
			conn.Write(GenerateIRCMessageBin(RplMyInfo, IRCUsername, fmt.Sprintf(":%s twIRC DOQRSZaghilopswz CFILMPQSbcefgijklmnopqrstvz bkloveqjfI", hostname)))
			conn.Write(GenerateIRCMessageBin(RplMotdStart, IRCUsername, ":Filling in a MOTD here because I have to."))
			conn.Write(GenerateIRCMessageBin(RplMotdEnd, IRCUsername, ":done"))
		}

		if strings.HasPrefix(line, "USER ") && ConnectionStage == 1 {
			if IRCUsername != "" {
				ConnectionStage++
			}
		}

		if strings.HasPrefix(line, "JOIN ##twitterstream") && ConnectionStage == 2 {
			conn.Write([]byte(fmt.Sprintf(":%s!~%s@idkwhatyourhostis JOIN ##twitterstream * :Ben Cox", IRCUsername, IRCUsername)))
			conn.Write(GenerateIRCMessageBin(RplNamReply, IRCUsername, fmt.Sprintf("@ ##twitterstream :@%s RandomGuy", IRCUsername)))
			conn.Write(GenerateIRCMessageBin(RplEndOfNames, IRCUsername, "##twitterstream :End of /NAMES list."))
		}

		if strings.HasPrefix(line, "MODE ##twitterstream") && ConnectionStage == 2 {
			conn.Write(GenerateIRCMessageBin(RplChannelModeIs, IRCUsername, "##twitterstream +ns"))
			conn.Write(GenerateIRCMessageBin(RplChannelCreated, IRCUsername, "##twitterstream 1401629312"))
			go StreamTwitter(conn, logindata, c)
			go PingClient(conn)
		}

	}

}

func PingClient(conn net.Conn) {
	for {
		_, e := conn.Write([]byte("PING :SUP\r\n"))
		if e != nil {
			break
		}
		time.Sleep(time.Second * 30)
	}
}

func StreamTwitter(conn net.Conn, logindata oauth.AccessToken, c *oauth.Consumer) {

	var response *http.Response
	response, e := c.Get(
		"https://userstream.twitter.com/1.1/user.json",
		map[string]string{},
		&logindata)

	if e != nil {
		return
	}

	twitterinbound := bufio.NewReader(response.Body)

	for {
		line, _, e := twitterinbound.ReadLine()

		if *logtwitter {
			ioutil.WriteFile("./debug_log", line, 744)
		}

		if e != nil {
			conn.Write([]byte(fmt.Sprintf(":SYS!~SYS@twitter.com PRIVMSG ##twitterstream : TWITTERSTREAM HAS BROKEN, HANGING UP. SORRY.\r\n")))
			conn.Close()
			return
		}
		var T Tweet
		e = json.Unmarshal(line, &T)
		if e == nil {
			TweetString := strings.TrimSpace(T.Text)
			TweetString = strings.Replace(TweetString, "\r", " ", -1)
			TweetString = strings.Replace(TweetString, "\n", " ", -1)
			conn.Write([]byte(fmt.Sprintf(":%s!~%s@twitter.com PRIVMSG ##twitterstream :%s\r\n", T.User.ScreenName, T.User.ScreenName, TweetString)))
		}
	}

	response.Body.Close()

}

func GenerateIRCMessage(code string, username string, data string) string {
	return fmt.Sprintf(":twitter.com %s %s %s\r\n", code, username, data)
}

func GenerateIRCMessageBin(code string, username string, data string) []byte {
	return []byte(GenerateIRCMessage(code, username, data))
}

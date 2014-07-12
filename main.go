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
	"strconv"
	"strings"
	"time"
)

var configarray []string

func main() {
	hostcfg := flag.String("listen", "localhost:6667", "<host>:<port>")
	flag.Parse()
	configbytes, err := ioutil.ReadFile("./twitterauth.cfg")
	if err != nil {
		log.Fatal("Could not read the config file. not going to bother")
	}

	configarray = strings.Split(strings.Replace(string(configbytes), "\r", "", -1), "\n")
	if len(configarray) != 2 && (len(configarray) != 3) {
		if !(len(configarray) == 3 && configarray[2] == "") {
			log.Fatal("Bad amount of data in config.")
		}
	}

	if configarray[0] == "API key" {
		log.Fatal("You need to fill out the config files...")
	}

	// Listen for incoming connections.
	l, err := net.Listen("tcp", *hostcfg)
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

var ReplyLatestTweet bool = false

// Handles incoming requests.
func handleIRCConn(conn net.Conn) {
	var ConnectionStage int = 0
	var TwitterToken string
	var IRCUsername string
	var LastTweetIDMap map[string]Tweet
	var LastMentionIDMap map[string]Tweet
	LastTweetIDMap = make(map[string]Tweet)
	LastMentionIDMap = make(map[string]Tweet)
	var HasAuthed bool = false
	var IsInChan bool = false

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
	var RQT *oauth.RequestToken
	reader := bufio.NewReader(conn)
	for {
		lineb, _, err := reader.ReadLine()
		line := string(lineb)

		if err != nil {
			return
		}

		fmt.Println(line)

		if strings.HasPrefix(line, "QUIT ") {
			conn.Close()
			return
		}

		if strings.HasPrefix(line, "PASS ") && ConnectionStage == 0 {
			TwitterToken = strings.Split(line, " ")[1]
			json.Unmarshal([]byte(TwitterToken), &logindata)
			fmt.Printf("Twitter token: %s \n", TwitterToken)
			ConnectionStage++
		}

		if strings.HasPrefix(line, "KICK ##twitterstream ") && HasAuthed && IsInChan {
			Target := strings.Split(line, " ")[2]
			r, e := c.Post("https://api.twitter.com/1.1/friendships/destroy.json",
				map[string]string{
					"screen_name": Target,
				},
				&logindata)
			ioutil.ReadAll(r.Body)
			if e == nil {
				conn.Write([]byte(fmt.Sprintf(":%s!~%s@twitter.com PART ##twitterstream :Unfollowed\r\n", Target, Target)))
			} else {
				conn.Write(GenerateIRCPrivateMessage("Unable to unfollow user.", "##twitterstream", "SYS"))
			}

		}

		if strings.HasPrefix(line, "NICK ") && ConnectionStage == 1 {
			fmt.Println(line)
			IRCUsername = strings.Split(line, " ")[1]
			conn.Write(GetWelcomePackets(IRCUsername, hostname))
			HasAuthed = true
		} else if strings.HasPrefix(line, "NICK ") && ConnectionStage == 0 {
			IRCUsername = strings.Split(line, " ")[1]
			conn.Write(GetWelcomePackets(IRCUsername, hostname))

			var url string
			var err error
			RQT, url, err = c.GetRequestTokenAndUrl("oob")
			if err != nil {
				log.Fatal(err)
			}

			conn.Write(GenerateIRCPrivateMessage(fmt.Sprintf("(1) Go to: %s", url), IRCUsername, "SYS"))
			conn.Write(GenerateIRCPrivateMessage("(2) Grant access, you should get back a verification code.", IRCUsername, "SYS"))
			conn.Write(GenerateIRCPrivateMessage("(3) Please enter the code as a raw command, EG '/347527'", IRCUsername, "SYS"))
			ConnectionStage = -1
		}

		// try and parse the string as a number to see what would happen
		linen := strings.TrimSpace(string(lineb))
		_, err = strconv.ParseInt(linen, 10, 64)
		if err == nil && ConnectionStage == -1 {

			accessToken, err := c.AuthorizeToken(RQT, linen)
			if err != nil {
				conn.Write(GenerateIRCPrivateMessage("Failed to auth you, Reconnect and try again (sry)", IRCUsername, "SYS"))
				return
			}
			conn.Write(GenerateIRCPrivateMessage("Okay next time you login use the PASS:", IRCUsername, "SYS"))
			b, _ := json.Marshal(accessToken)
			conn.Write(GenerateIRCPrivateMessage(fmt.Sprintf("%s", string(b)), IRCUsername, "SYS"))
			conn.Write(GenerateIRCPrivateMessage("You MUST reconnect now since all commands have been disabled.", IRCUsername, "SYS"))
			return
		}

		if strings.HasPrefix(line, "USER ") && ConnectionStage == 1 {
			if IRCUsername != "" {
				ConnectionStage++
			}
		}

		if strings.HasPrefix(strings.ToUpper(line), "MENTION") && HasAuthed && IsInChan {
			ReplyLatestTweet = false
			conn.Write(GenerateIRCPrivateMessage("PM's will now RE the latest mention of you", IRCUsername, "SYS"))
		}

		if strings.HasPrefix(strings.ToUpper(line), "ALL") && HasAuthed && IsInChan {
			ReplyLatestTweet = true
			conn.Write(GenerateIRCPrivateMessage("PM's will now RE the latest tweet of the target", IRCUsername, "SYS"))
		}

		if strings.HasPrefix(strings.ToUpper(line), "UNDO") && HasAuthed && IsInChan {
			_, err := c.Post(
				fmt.Sprintf("https://api.twitter.com/1.1/statuses/destroy/%s.json", LastTweetIDMap[strings.ToLower(IRCUsername)].IdStr),
				map[string]string{
					"id": LastTweetIDMap[strings.ToLower(IRCUsername)].IdStr,
				},
				&logindata)
			if err != nil {
				conn.Write(GenerateIRCPrivateMessage("Failed to undo tweet. ur boned now.", "##twitterstream", "SYS"))
			} else {
				conn.Write(GenerateIRCPrivateMessage("Tweet undone", IRCUsername, "SYS"))
			}
		}

		if strings.HasPrefix(line, "JOIN ##twitterstream") && HasAuthed {
			conn.Write([]byte(fmt.Sprintf(":%s!~%s@twitter.com JOIN ##twitterstream * :Ben Cox\r\n", IRCUsername, IRCUsername)))
			NList := ProduceNameList(logindata, c)
			for _, v := range NList {
				conn.Write(GenerateIRCMessageBin(RplNamReply, IRCUsername, fmt.Sprintf("@ ##twitterstream :@%s %s", IRCUsername, v)))
			}
			conn.Write(GenerateIRCMessageBin(RplEndOfNames, IRCUsername, "##twitterstream :End of /NAMES list."))
			go StreamTwitter(conn, logindata, c, LastTweetIDMap, LastMentionIDMap, IRCUsername)
			go PingClient(conn)
			IsInChan = true
		}

		if strings.HasPrefix(line, "MODE ##twitterstream") && HasAuthed {
			conn.Write(GenerateIRCMessageBin(RplChannelModeIs, IRCUsername, "##twitterstream +ns"))
			conn.Write(GenerateIRCMessageBin(RplChannelCreated, IRCUsername, "##twitterstream 1401629312"))
		}
		// PRIVMSG ##twitterstream :Holla
		if strings.HasPrefix(line, "PRIVMSG ##twitterstream :") && HasAuthed && IsInChan {
			_, err := c.Post(
				"https://api.twitter.com/1.1/statuses/update.json",
				map[string]string{
					"status": strings.Replace(line, "PRIVMSG ##twitterstream :", "", 1),
				},
				&logindata)
			if err != nil {
				conn.Write(GenerateIRCPrivateMessage("Failed to post tweet.", "##twitterstream", "SYS"))
			}
		} else if strings.HasPrefix(line, "PRIVMSG ") && HasAuthed {
			HandlePRIVreply(conn, logindata, c, line, LastTweetIDMap, LastMentionIDMap)
		}

	}

}

func HandlePRIVreply(conn net.Conn, logindata oauth.AccessToken, c *oauth.Consumer, line string, LastTweetIDMap map[string]Tweet, LastMentionIDMap map[string]Tweet) {
	bits := strings.Split(line, " ")
	if len(bits) > 2 {
		tweetstring := strings.Replace(line, "PRIVMSG "+bits[1], "", 1)
		var err error
		var lastmention Tweet
		if ReplyLatestTweet {
			lastmention = LastTweetIDMap[strings.ToLower(bits[1])]
		} else {
			lastmention = LastMentionIDMap[strings.ToLower(bits[1])]
		}

		if lastmention.User.IdStr != "" {
			_, err = c.Post(
				"https://api.twitter.com/1.1/statuses/update.json",
				map[string]string{
					"status":                fmt.Sprintf("@%s %s", bits[1], tweetstring[2:]),
					"in_reply_to_status_id": fmt.Sprint(lastmention),
				},
				&logindata)
			fmt.Printf("I'm going to post '%s' with a msg ID chain %s %s \n", "@"+bits[1]+" "+tweetstring, lastmention, fmt.Sprint(lastmention))
		} else {
			_, err = c.Post(
				"https://api.twitter.com/1.1/statuses/update.json",
				map[string]string{
					"status": fmt.Sprintf("@%s %s", bits[1], tweetstring[2:]),
				},
				&logindata)
			fmt.Printf("I'm going to post '%s' \n", "@"+bits[1]+" "+tweetstring)
		}
		if err != nil {
			conn.Write(GenerateIRCPrivateMessage("Failed to post tweet.", "##twitterstream", "SYS"))
		}
	}
}

func PingClient(conn net.Conn) {
	for {
		_, e := conn.Write([]byte(fmt.Sprintf("PING :%d\r\n", int32(time.Now().Unix()))))
		if e != nil {
			break
		}
		time.Sleep(time.Second * 30)
	}
}

func StreamTwitter(conn net.Conn, logindata oauth.AccessToken, c *oauth.Consumer,
	LastTweetIDMap map[string]Tweet, LastMentionIDMap map[string]Tweet, username string) {

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

		if e != nil {
			conn.Write(GenerateIRCPrivateMessage("System has broken, Shutting down.", "##twitterstream", "SYS"))
			conn.Close()
			return
		}
		var T Tweet
		e = json.Unmarshal(line, &T)
		if e == nil && T.Text != "" {
			LastTweetIDMap[strings.ToLower(T.User.ScreenName)] = T
			TweetString := strings.TrimSpace(T.Text)
			TweetString = strings.Replace(TweetString, "\r", " ", -1)
			TweetString = strings.Replace(TweetString, "\n", " ", -1)
			conn.Write(GenerateIRCPrivateMessage(TweetString, "##twitterstream", T.User.ScreenName))
			if strings.HasPrefix(strings.ToLower(T.Text), "@"+strings.ToLower(username)) {
				LastMentionIDMap[strings.ToLower(T.User.ScreenName)] = T
				conn.Write(GenerateIRCPrivateMessage(TweetString, username, T.User.ScreenName))
			}
		} else if T.Text == "" && e == nil && !strings.Contains(string(line), "friends") {
			// Maybe its a delete packet???
			DP := RemovePacket{}
			e = json.Unmarshal(line, &DP)

			if DP.Delete.Status.IdStr != "" {
				OffendersName := ScanForName(DP.Delete.Status.UserIdStr, LastTweetIDMap)
				ExtraInfo := ""
				if LastTweetIDMap[OffendersName].IdStr != "" {
					if LastTweetIDMap[OffendersName].IdStr == DP.Delete.Status.IdStr {
						ExtraInfo = "(It was the latest tweet they posted)"
					}
				}
				conn.Write(GenerateIRCPrivateMessage(fmt.Sprintf("User %s (%s) remove tweet %s %s", DP.Delete.Status.UserIdStr, OffendersName, DP.Delete.Status.IdStr, ExtraInfo), "##twitterstream", "SYS"))
			} else {
				conn.Write(GenerateIRCPrivateMessage("unknown message: "+string(line), "##twitterstream", "SYS"))
			}
		}
	}

	response.Body.Close()

}

func ScanForName(ID string, UserMap map[string]Tweet) string {
	for k, v := range UserMap {
		if ID == v.User.IdStr {
			return k
		}
	}
	return "???"
}

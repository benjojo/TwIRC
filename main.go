package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "6667"
	CONN_TYPE = "tcp"
)

var configarray []string

func main() {
	configbytes, err := ioutil.ReadFile("./twitterauth.cfg")
	if err != nil {
		log.Fatal("Could not read the config file. not going to bother")
	}

	configarray = strings.Split(strings.Replace(string(configbytes), "\r", "", -1), "\n")
	if len(configarray) != 2 {
		log.Fatal("bad amount of data in config.")
	}

	if configarray[0] == "API key" {
		log.Fatal("you need to fill out the config files...")
	}

	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
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

		if strings.HasPrefix(line, "PASS ") && ConnectionStage == 0 {
			fmt.Println(line)
			TwitterToken = strings.Split(line, " ")[1]
			json.Unmarshal([]byte(TwitterToken), &logindata)
			fmt.Printf("Twitter token: %s \n", TwitterToken)
			ConnectionStage++
			// var responce http.Response
			// response, _ := c.Get(
			// 	"https://userstream.twitter.com/1.1/user.json",
			// 	map[string]string{},
			// 	&logindata)

			// twitterinbound := bufio.NewReader(responce.Body)
			// conn.Write(b)
			// response.Body.Close()
		}

		if strings.HasPrefix(line, "NICK ") && ConnectionStage == 1 {
			fmt.Println(line)
			IRCUsername := strings.Split(line, " ")[1]
		}

	}

}

func GenerateIRCMessage(code string, username string, data string) string {
	return fmt.Sprintf(":twitter.com %s %s :%s", code, username, data)
}

func GenerateIRCMessageBin(code string, username string, data string) string {
	return []byte(GenerateIRCMessage(code, username, data))
}

package main

import (
	"fmt"
)

func GenerateIRCMessage(code string, username string, data string) string {
	return fmt.Sprintf(":twitter.com %s %s %s\r\n", code, username, data)
}

func GenerateIRCMessageBin(code string, username string, data string) []byte {
	return []byte(GenerateIRCMessage(code, username, data))
}

func GetWelcomePackets(IRCUsername string, hostname string) []byte {
	pack := ""
	pack += GenerateIRCMessage(RplWelcome, IRCUsername, ":Welcome to TwiRC")
	pack += GenerateIRCMessage(RplYourHost, IRCUsername, fmt.Sprintf(":Host is: %s", hostname))
	pack += GenerateIRCMessage(RplCreated, IRCUsername, ":This server was first made on 31/06/2014")
	pack += GenerateIRCMessage(RplMyInfo, IRCUsername, fmt.Sprintf(":%s twIRC DOQRSZaghilopswz CFILMPQSbcefgijklmnopqrstvz bkloveqjfI", hostname))
	pack += GenerateIRCMessage(RplMotdStart, IRCUsername, ":Filling in a MOTD here because I have to.")
	pack += GenerateIRCMessage(RplMotdEnd, IRCUsername, ":done")
	return []byte(pack)
}

func GenerateIRCPrivateMessage(content string, room string, username string) []byte {
	return []byte(fmt.Sprintf(":%s!~%s@twitter.com PRIVMSG %s :%s\r\n", username, username, room, content))
}

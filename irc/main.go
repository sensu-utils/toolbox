package main

import (
	"crypto/tls"
	"log"
	"net"
	"os"

	"github.com/belak/sensu-go-tools/utils"
	"github.com/go-irc/irc"
)

type Config struct {
	Nick     string `json:"irc_nick"`
	Server   string `json:"irc_server"`
	Password string `json:"irc_password"`
	SSL      bool   `json:"irc_ssl"`
	Channel  string `json:"irc_channel"`
}

var config Config
var event utils.Event

func main() {
	utils.InitPlugin("irc", &event, &config)

	var err error

	log.Println(event)

	var rawConn net.Conn
	if config.SSL {
		rawConn, err = tls.Dial("tcp", config.Server, nil)
	} else {
		rawConn, err = net.Dial("tcp", config.Server)
	}
	if err != nil {
		log.Fatalln(err)
	}

	conn := irc.NewConn(rawConn)
	if config.Password != "" {
		conn.Writef("PASS :%s", config.Password)
	}
	conn.Writef("NICK :%s", config.Nick)
	conn.Writef("USER %s 0.0.0.0 0.0.0.0 :%s", "sensu", "sensu")

	for {
		msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatalln(err)
		}

		if msg.Command == "PING" {
			reply := msg.Copy()
			reply.Command = "PONG"
			conn.WriteMessage(reply)
		} else if msg.Command == "001" {
			conn.Writef("JOIN :%s", config.Channel)
			conn.Writef("PRIVMSG %s :%s", config.Channel, "Message")
			conn.Writef("QUIT :bye")

			os.Exit(0)
			return
		}
	}
}

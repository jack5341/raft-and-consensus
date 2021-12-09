package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/jack5341/raft-consensus/src/service/cmd/handle"
)

const (
	CONN_HOST = "localhost"
	CONN_TYPE = "tcp"
)

type AppendEntries struct {
	term         string
	leaderId     int32
	prevLogIndex int32
	prevLogTerm  string
	entries      string
	leaderCommit string
}

type RequestVote struct {
	term         string
	candidateId  int32
	prevLogIndex int32
	prevLogTerm  string
}

func main() {
	port := flag.String("port", "", "port to listen on")
	ports := flag.String("targetPorts", "", "port to listen on")
	flag.Parse()

	CONN_ADRESS := ":" + *port

	if *ports != "" {
		splitedString := strings.Split(*ports, ",")

		for _, portServer := range splitedString {
			c, err := net.Dial(CONN_TYPE, "localhost:"+portServer)

			if err != nil {
				log.Fatal(err)
			}

			for {
				data, err := c.Write([]byte("Hello"))

				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("data:", data)
				time.Sleep(time.Millisecond * 3000)
			}
		}
	}

	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+CONN_ADRESS)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// Handle connections in a new goroutine.
		handle.Request(conn)
	}
}

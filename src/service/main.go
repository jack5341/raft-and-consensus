package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
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

		// print hello world to rbrowser.
		for _, portServer := range splitedString {
			c, err := net.Dial(CONN_TYPE, "localhost:"+portServer)

			if err != nil {
				log.Fatal(err)
			}

			for {
				nodeID := rand.Intn(100)
				entrie := AppendEntries{
					term:         "1",
					leaderId:     int32(nodeID),
					prevLogIndex: 0,
					prevLogTerm:  "1",
					leaderCommit: "1,2,3,4,5",
				}

				rand.Seed(time.Now().UnixNano())
				n := rand.Intn(5)

				_, err := c.Write([]byte(fmt.Sprintf("%v \n", entrie)))

				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Sleeping %d seconds...\n", n)
				time.Sleep(time.Duration(n) * time.Second)
			}
		}
	}

	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+CONN_ADRESS)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Listen for an incoming connection.
	conn, err := l.Accept()

	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}

	// Handle connections in a new goroutine.
	handle.Request(conn)
}

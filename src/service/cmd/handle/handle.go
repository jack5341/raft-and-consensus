package handle

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type AppendEntries struct {
	term         string
	leaderId     int32
	prevLogIndex int32
	prevLogTerm  string
	entries      string
	leaderCommit string
}

// Handles incoming requests.
func Request(conn net.Conn) {
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(data)

		conn.Write([]byte("Message received."))

		return
	}
}

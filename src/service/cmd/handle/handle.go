package handle

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Handles incoming requests.
func Request(conn net.Conn) {
	for {
		fmt.Println(conn)
		input, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(input)
		conn.Write([]byte("Message received."))
	}
}

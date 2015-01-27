package main

import (
	"log"
	"net"
	"strconv"
)

func LocalApi() {

	// Set up listener for defined host and port
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(8881))

	if err != nil {
		log.Println(err)
	}

	// set up channel on which to send accepted connections
	listen := make(chan net.Conn, 100)
	go acceptConnection(listener, listen)

	for {
		select {
		case conn := <-listen:
			go handleClient(conn)
		}
	}
}

// Accept a client connection and collect it in a channel
func acceptConnection(listener net.Listener, listen chan<- net.Conn) {
	log.Println("ACCEPT CONNECTION")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		listen <- conn
	}
}

func handleClient(client net.Conn) {
	log.Println("HANDLE CLIENT")
	for {
		buf := make([]byte, 4096)
		numbytes, err := client.Read(buf)
		if numbytes == 0 || err != nil {
			return
		}
		log.Println(string(buf))
		client.Write(buf)
	}
}

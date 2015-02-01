package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var stdlog, errlog *log.Logger

func Storage() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Set up listener for defined host and port
	listener, err := net.Listen("tcp", ":9977")
	if err != nil {
		//return "Possibly was a problem with the port binding", err
		fmt.Println("Possibly was a problem with the port binding")
	}

	// set up channel on which to send accepted connections
	listen := make(chan net.Conn, 100)
	go acceptConnection(listener, listen)

	// loop work cycle with accept connections or interrupt
	// by system signal
	for {
		select {
		case conn := <-listen:
			go handleClient(conn)
		case killSignal := <-interrupt:
			stdlog.Println("Got signal:", killSignal)
			stdlog.Println("Stoping listening on ", listener.Addr())
			listener.Close()
			if killSignal == os.Interrupt {
				fmt.Println("Daemon was interruped by system signal")
				//return "Daemon was interruped by system signal", nil
			}
			fmt.Println("Daemon was killed")
			//return "Daemon was killed", nil
		}
	}
}
func acceptConnection(listener net.Listener, listen chan<- net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		listen <- conn
	}
}

func handleClient(client net.Conn) {
	for {
		fmt.Println("log")
		// buf := make([]byte, 4096)
		// numbytes, err := client.Read(buf)
		// if numbytes == 0 || err != nil {
		// 	return
		// }
		// client.Write(buf)
	}
}

/**
 * Penlook Project
 *
 * Copyright (c) 2015 Penlook Development Team
 *
 * --------------------------------------------------------------------
 *
 * This program is free software: you can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License
 * as published by the Free Software Foundation, either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public
 * License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 *
 * --------------------------------------------------------------------
 *
 * Author:
 *     Loi Nguyen       <loint@penlook.com>
 *     Tin Nguyen       <tinntt@penlook.com>
 */

package main

import (
	"fmt"
	. "github.com/penlook/daemon"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	service := Service{
		Process: Storage,
	}

	service.GetInfo("storage")
	service.Initialize()
}

func Storage(service Service) {
	s3 := S3{}
	s3.GetConfig()
	s3.Create()

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
			go handleClient(conn, s3)
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

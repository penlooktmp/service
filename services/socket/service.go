// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package main

import (
	"log"
    "net/http"
    "os/exec"
    "fmt"
    "github.com/googollee/go-socket.io"
	"github.com/penlook/daemon"
)

func main() {

	service := daemon.Service{
		Name:        "socket",
		Description: "Penlook Socket Service",
		Process:     Socket,
	}

	service.Initialize()
}

func Socket() {
	server, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    server.On("connection", func(so socketio.Socket) {
        log.Println("on connection")
        so.Join("chat")
        so.On("chat message", func(msg string) {
            log.Println("emit:", so.Emit("chat message", msg))
            so.BroadcastTo("chat", "chat message", msg)
        })
        so.On("disconnection", func() {
            log.Println("on disconnect")
        })
    })
    server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })

    str, err := exec.Command("pwd").Output();
    if err != nil {
    	panic(err)
    }
    fmt.Printf("in all caps: %q\n", str.String()); 
    http.Handle("/socket.io/", server)

    http.Handle("/abc", http.FileServer(http.Dir("./asset")))
    log.Println("Serving at localhost:5000...")
    log.Fatal(http.ListenAndServe(":5000", nil))
}

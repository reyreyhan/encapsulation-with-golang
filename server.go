package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type WinLose struct {
	Hero_id int32
	WinLose bool
}

func main() {
	fmt.Println("Server")
	// start TCP server and listen on port 8000
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		// handle error
		panic(err)
	}

	// this blocks until connection or error
	conn, err := ln.Accept()
	if err != nil {
		// handle error
		fmt.Println("Error")
	}
	// create new decoder object and provide connection
	dec := gob.NewDecoder(conn)
	// create blank object
	p := &WinLose{}
	// decode serialize data
	dec.Decode(p)
	// print
	fmt.Println("Hero ID : ", p.Hero_id, " - ", p.WinLose)
	// close connection for that client
	conn.Close()
}

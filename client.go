package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type WinLose struct {
	Hero_id int32
	WinLose bool
}

func main() {
	fmt.Println("Client")
	//create structure object
	winLoseEncode := WinLose{Hero_id: 23, WinLose: false}

	fmt.Println("start client")
	// dial TCP connection
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	//Create encoder object, We are passing connection object in Encoder
	encoder := gob.NewEncoder(conn)
	// Encode Structure, IT will pass object over TCP connection
	encoder.Encode(winLoseEncode)
	// close connection
	conn.Close()
	fmt.Println("done")
}

package main

import (
	"encoding/json"
	"fmt"
	"net"
)

var msg request

func main() {
	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		fmt.Println(err)
		panic("[ERROR] Can't listen the network")
	}
	fmt.Println("Launching server")
	for {
		conn, errcon := listener.AcceptTCP()
		if errcon != nil {
			fmt.Println(errcon)
			panic("[ERROR] Nothing is sent to the network")
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) {
	decoder := json.NewDecoder(conn)
	for decoder.More() {
		err := decoder.Decode(&msg)
		if err != nil {
			panic(err)
		}
		fmt.Println("Message recieved:", msg)
		encoder := json.NewEncoder(conn)
		n, d := getNth(msg.Number)
		fmt.Println("Response sent:", n, d)
		err = encoder.Encode(response{n, d})
		if err != nil {
			panic(err)
		}
	}
	defer conn.Close()
}

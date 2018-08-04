package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

var ans response

func main() {
	var msg request
	conn, err := net.DialTCP("tcp", nil, laddr)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		encoder := json.NewEncoder(conn)
		str := scanner.Text()
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
			panic("[ERROR] Can't convert to int")
		}
		msg = request{num}
		if conn != nil {
			encoder.Encode(&msg)
		}
		go RecieveResponse(conn)
	}
	defer conn.Close()
}

func RecieveResponse(conn *net.TCPConn) {
	decoder := json.NewDecoder(conn)
	for decoder.More() {
		decoder.Decode(&ans)
		fmt.Println(ans.Dur, ans.Num)
	}
}

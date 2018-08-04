package main

import (
	"math/big"
	"net"
	"time"
)

type request struct {
	Number int
}

type response struct {
	Num *big.Int
	Dur time.Duration
}

var ip4 net.IP = net.IPv4(127, 0, 0, 1)
var port int = 8080
var laddr *net.TCPAddr = &net.TCPAddr{ip4, port, ""}

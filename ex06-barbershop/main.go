package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}
var wg sync.WaitGroup
var customers = [6]Customer{}
var next chan Customer = make(chan Customer)
var count int = 0

func main() {
	fmt.Println("<<<<Barber has opened the shop!>>>>")
	mybarber := Barber{true}
	wg.Add(7)
	go mybarber.IsWorking()
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			time.Sleep(1 * time.Second)
		}
		go NewCustomer()
	}
	wg.Wait()
	close(next)
}

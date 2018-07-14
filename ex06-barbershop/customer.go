package main

import (
	"fmt"
	"time"
)

type Customer struct {
	occupied bool
	seatnum  int
}

// func NewCustomer(): Creates a customer and puts him in the queue if there is space otherwise the customer leaves
func NewCustomer() {
	count++

	fmt.Printf("Customer # %d comes into the shop\n", count)
	for r := range customers {
		if customers[r].occupied == false {
			customers[r].occupied = true
			customers[r].seatnum = r
			next <- customers[0]
			//fmt.Println(customers)
			customers[0].seatnum = 0
			customers[0].occupied = false
			ShiftCustomers()
			wg.Done()
			return
		}
	}
	fmt.Printf("Customer %d is leaving, because there are way too many people...\n", count)
	time.Sleep(2 * time.Second)
	return
}

//func ShiftCustomers(): Shift customers towards the beginning so that they are served in the right order
func ShiftCustomers() {
	for i := 0; i < len(customers)-1; i++ {
		customers[i] = customers[i+1]
		customers[i+1].seatnum = 0
		customers[i+1].occupied = false
	}
}

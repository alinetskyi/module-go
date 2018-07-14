package main

import (
	"fmt"
	"time"
)

type Barber struct {
	sleeps bool
}

//func (barber *Barber) IsWorking() : Checks whether there is someone in the queue and cuts his hair, otherwise is sleeping
func (barber *Barber) IsWorking() {
	var cs int = 1
	for {
		select {
		case n, ok := <-next:
			if ok {
				if n.occupied {
					barber.sleeps = false
					fmt.Printf("The Barber gives a haircut to customer %d\n", cs)
					cs++
					time.Sleep(4 * time.Second)
				} else {
					barber.sleeps = true
					fmt.Println("Barber is sleeping..zzzzZZZZzzzz")
					time.Sleep(2 * time.Second)
				}
			} else {
				fmt.Println("Barbershop is closed")
			}
		default:
			barber.sleeps = true
			fmt.Println("Barber is sleeping..zzzzZZZZzzzz")
			time.Sleep(2 * time.Second)

		}
	}
	//fmt.Println("Something went wrong")
	wg.Done()
}

package main

import (
	"math/big"
	"time"
)

func fib() func() *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	return func() *big.Int {
		a.Add(a, b)
		a, b = b, a
		return a
	}
}

func getNth(num int) (*big.Int, time.Duration) {
	start := time.Now()
	var res *big.Int
	f := fib()
	for i := 0; i < num; i++ {
		res = f()
	}
	end := time.Now()
	return res, end.Sub(start)
}

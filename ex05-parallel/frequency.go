package letter

import (
	"sync"
)

//Defining global varaibles
var wg sync.WaitGroup
var mymap map[string]int = make(map[string]int)
var mutex sync.Mutex
var flag bool = false
var pos int = 0

// func ConcurrentFrequency(str []string) map[string]int : returns map with characters        and their frequencies without concurency
func Frequency(str string) map[string]int {
	for i := pos; i < len(str); i++ {
		mutex.Lock()
		mymap[string(str[pos])]++
		pos++
		mutex.Unlock()
	}
	if flag == true {
		defer wg.Done()
	}
	return mymap
}

// func ConcurrentFrequency(str []string) map[string]int : returns map with characters and their frequencies using concurencies
func ConcurrentFrequency(str []string) map[string]int {
	wg.Add(len(str))
	flag = true
	for i := 0; i < len(str); i++ {
		go Frequency(str[i])
	}
	wg.Wait()
	return mymap
}

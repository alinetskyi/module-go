package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, jobs <-chan string) {
	var spawned bool
	defer wg.Done()
	for j := range jobs {
		if spawned != true {
			fmt.Printf("worker:%d spawning\n", id)
		}
		spawned = true
		l, _ := time.ParseDuration(j + "s")
		fmt.Printf("worker:%d sleep:%s\n", id, j)
		time.Sleep(l)
	}
	if spawned == true {
		fmt.Printf("worker:%d stopping\n", id)
	}
}

func Run(num int) {
	jobs := make(chan string, 256)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var k int = 1
	for scanner.Scan() {
		if k <= num {
			wg.Add(1)
			go worker(k, jobs)
			k++
		}
		txt := scanner.Text()
		jobs <- txt
	}
	close(jobs)
	wg.Wait()
}

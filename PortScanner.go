// Port Scanner in GO

package main

import (
	"fmt"
	"net"
	"os"
	"sort"
)

func worker(ports, results chan int) {
	server := os.Args[1]
	for p := range ports {
		address := fmt.Sprintf("%s:%d", server, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 500)
	results := make(chan int)
	var openports []int
	// var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
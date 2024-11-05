package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

var numServers int

func main() {
	// Set up logging to a file
	ArgunemtHandler()

	file, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Could not open log file:", err)
		return
	}
	defer file.Close()
	log.SetOutput(file)

	var wg sync.WaitGroup

	for i := 0; i < numServers; i++ {
		wg.Add(1)
		go startServer(&wg)
	}

	// Load balancer setup
	lb := createLoadBalancer()
	go func() {
		if err := lb.Run(":3000"); err != nil {
			log.Println("Load balancer failed:", err)
		}
	}()

	wg.Wait()
}

func ArgunemtHandler() {
	if len(os.Args) < 2 {
		log.Println("Usage: go run . <num of server>")
		return
	}

	migrator := os.Args[1]

	numServer, err := strconv.Atoi(migrator)
	if err != nil {
		fmt.Println("Please use num")
		return
	}

	if migrator != "" {
		numServers = numServer
	} else {
		log.Println("Usage: go run . <env_file> [create/migrate/seed (optional)]")
	}
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

var numServers int
var stringServers []int
var start bool

func main() {
	ArgumentHandler()
}

func ArgumentHandler() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: go run . <num of server>")
		return
	}

	for i, arg := range os.Args[1:] {
		if i%2 == 0 {
			switch arg {
			case "-num":

				numServer, err := strconv.Atoi(os.Args[i+2])
				fmt.Println(numServer)
				if err != nil {
					fmt.Println("Please use valid integer number")
					return
				}
				numServers = numServer

			default:
				fmt.Println("Numserver = ", numServers)
			}
		}
	}

	Run()
}

func Run() {
	file, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Could not open log file:", err)
		return
	}
	defer file.Close()
	log.SetOutput(file)

	var wg sync.WaitGroup

	if numServers > 0 {
		fmt.Println("masuk num")
		for i := 0; i < numServers; i++ {
			wg.Add(1)
			port := rand.Intn(10000) + 1000
			go startServer(&wg, port)
		}
	} else if numServers == 0 {
		for i, _ := range stringServers[0:] {
			wg.Add(1)
			if start {
				go startServer(&wg, stringServers[i])
			}
		}
	}

	lb := createLoadBalancer()
	go func() {
		if err := lb.Run(":3000"); err != nil {
			log.Println("Load balancer failed:", err)
		}
	}()

	wg.Wait()
}

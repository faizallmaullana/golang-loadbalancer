package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"golang.org/toolchain/src/math/rand"
)

// Start the server on the specified port.
func startServer(wg *sync.WaitGroup, port int) {
	defer wg.Done()

	// Get mode from environment variables
	mode := os.Getenv("MODE")
	gin.SetMode(mode)
	r := gin.New()

	rand.Seed(time.Now().UnixNano())
	address := fmt.Sprintf("http://localhost:%d", port)

	mu.Lock()
	servers = append(servers, address)
	mu.Unlock()

	r.Use(static.Serve("/", static.LocalFile("./dist", true)))

	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	// Start the server
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Printf("Server failed on port %d: %v", port, err)
	}
}

package main

import (
	"io"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var servers []string
var mu sync.Mutex
var nextServer int

func createLoadBalancer() *gin.Engine {
	lb := gin.Default()
	lb.Any("/*any", loadBalancer)
	return lb
}

func loadBalancer(c *gin.Context) {
	mu.Lock()
	server := servers[nextServer]
	nextServer = (nextServer + 1) % len(servers)
	mu.Unlock()

	req, err := http.NewRequest(c.Request.Method, server+c.Request.RequestURI, c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to create request")
		return
	}

	for key, value := range c.Request.Header {
		req.Header.Set(key, value[0])
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.String(http.StatusServiceUnavailable, "Service Unavailable")
		return
	}
	defer resp.Body.Close()

	for key, value := range resp.Header {
		c.Header(key, value[0])
	}
	c.Status(resp.StatusCode)

	if _, err := io.Copy(c.Writer, resp.Body); err != nil {
		c.String(http.StatusInternalServerError, "Failed to read response")
	}
}

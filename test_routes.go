package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func testRoute(url string) {
	fmt.Printf("Testing %s...\n", url)
	
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response: %s\n\n", string(body))
}

func main() {
	baseURL := "http://localhost:8002"
	
	fmt.Println("Testing Event Service Routes...")
	fmt.Println("================================")
	
	// Test root route
	testRoute(baseURL + "/")
	
	// Test health route
	testRoute(baseURL + "/health")
	
	// Test existing API route
	testRoute(baseURL + "/api/v1/events")
}
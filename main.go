package main

import (
	_ "apitester/analytics"
	_ "apitester/charting"
	"apitester/config"
	"fmt"
)

// Read configuration file
// Initialize variables in step
// Execute calls in sequence or parallel
// Record measurements
// Clean and aggregate data
// Create graphs

func main() {
	config, _ := config.ReadConfig()
	fmt.Println(config.URL)
}

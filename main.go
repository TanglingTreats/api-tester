package main

import (
	_ "apitester/analytics"
	_ "apitester/charting"
	"apitester/config"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Read configuration file
// Initialize variables in step
// Execute calls in sequence or parallel
// Record measurements
// Clean and aggregate data
// Create graphs

// Unit of measurements
type Unit struct {
	value int
	unit  string
}

type Test struct {
	stepType config.StepType
	steps    []config.Step
}

func main() {
	// Get configuration
	config, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Error occurred, exiting")
		os.Exit(1)
	}

	_ = config.URL

	// Resolve steps
	tests := resolveSteps(config)

	// Execute API Tests
	stressTestAPI(config.URL, config.Token, tests)
}

func stressTestAPI(url string, token string, tests []Test) {
	client := &http.Client{}

	for _, test := range tests {
		for _, step := range test.steps {
			fmt.Println("\nExecuting: " + step.Name)
			req, err := http.NewRequest(string(step.Method), url+step.Endpoint, nil)
			if err != nil {
				fmt.Println(err)
			}
			req.Header.Add("Authorization", "Bearer "+token)
			req.Header.Add("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}

			data, err := io.ReadAll(resp.Body)
			fmt.Printf("%v\n", string(data))
		}
	}
}

func resolveSteps(conf config.Config) []Test {
	tests := []Test{}
	currStep := []config.Step{}

	// Loop through configuration steps
	stepType := config.Sequential
	for _, step := range conf.Steps {
		currStep = append(currStep, step)
		stepType = step.StepType
	}
	test := Test{stepType: stepType, steps: currStep}
	tests = append(tests, test)

	return tests
}

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

// Unit of measurements
type Unit struct {
	value int
	unit  string
}

type Test struct {
	stepType config.StepType
	step     []config.Step
}

func main() {
	// Get configuration
	config, _ := config.ReadConfig()
	fmt.Println(config.URL)
	fmt.Println(config.Version)

	_ = config.URL

	// Resolve steps
	tests := resolveSteps(config)
	fmt.Println(tests)

	// Execute steps
}

func resolveSteps(conf config.Config) []Test {
	tests := []Test{}
	currStep := []config.Step{}

	// Loop through configuration steps
	stepType := config.Sequential
	for _, step := range conf.Steps {
		fmt.Println(step)
		currStep = append(currStep, step)
		stepType = step.StepType
	}
	test := Test{stepType: stepType, step: currStep}
	tests = append(tests, test)

	return tests
}

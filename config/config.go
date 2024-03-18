package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Measurement int

const (
	ResponseTime Measurement = iota
	BlockingDuration
)

type TestType int

const (
	Sequence TestType = iota
	Parallel
)

type Step struct {
	Name      string      `json:"name"`
	Endpoint  string      `json:"endpoint"`
	TestType  TestType    `json:"testType"`
	Load      int         `json:"load"`
	DependsOn []string    `json:"dependsOn"`
	Measure   Measurement `json:"measure"`
}

type Config struct {
	Version int    `json:"version"`
	URL     string `json:"url"`
	Steps   []Step `json:"steps"`
}

func readConfig() ([]Step, error) {
	var steps []Step

	jsonFile, error := os.Open("configuration.json")

	if error != nil {
		fmt.Println(error)
		return steps, error
	}

	bytes, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(bytes, &steps)

	return steps, nil

}

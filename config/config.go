package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Step struct {
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

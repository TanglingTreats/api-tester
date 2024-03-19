package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadConfig() (Config, error) {
	var config Config

	jsonFile, error := os.Open("configuration.json")

	if error != nil {
		fmt.Println(error)
		return config, error
	}

	defer jsonFile.Close()

	b, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(b, &config)

	return config, nil

}

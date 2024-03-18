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

	bytes, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(bytes, &config)

	return config, nil

}

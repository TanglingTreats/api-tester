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

	err := json.Unmarshal(b, &config)

	if err != nil {
		fmt.Println("Error occurred unmarshalling")
		fmt.Println(err.Error())
		return config, error
	}

	return config, nil

}

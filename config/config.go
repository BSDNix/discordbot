package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	//public vars
	Token string
	BotPrefix string

	//private vars
	config *configStruct
)

type configStruct struct { 
	Token string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}
func ReadConfig() error {
	fmt.Println("Reading conf..")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil

} 

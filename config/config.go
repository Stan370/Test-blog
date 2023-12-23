package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Endpoint string `json:"db_endpoint"`
	DbName   string `json:"db_name"`
}

type Config struct {
	Database DatabaseConfig 
	//to be continued, app service configs
}

func loadConfig() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Panicln("Error reading config file:", err)
		return
	}

	var config Config
	// Unmarshal the JSON data into the Config struct
	if err := json.Unmarshal(data, &config); err != nil {
		log.Panicln("Error unmarshalling JSON:", err)
		return
	}

	log.Println("DB Name:", config.Database.DbName)
	fmt.Println("DB Endpoint:", config.Database.Endpoint)

}

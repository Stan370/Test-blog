package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Endpoint string `json:"db_endpoint"`
	DbName   string `json:"db_name"`
	//to be continued, app service configs
}

func (c *Config) GetDbConnection() string {
	return c.Username + ":" + c.Password + "@tcp(" + c.Endpoint + ")/" + c.DbName + "?parseTime=true"
}

func loadConfig() (*Config, error) {
	config := &Config{}
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Panicln("Error reading config file:", err)
		return nil, err
	}

	// Unmarshal the JSON data into the Config struct
	if err := json.Unmarshal(data, &config); err != nil {
		log.Panicln("Error unmarshalling JSON:", err)
		return nil, err
	}

	log.Println("DB Name:", config.DbName)
	fmt.Println("DB Endpoint:", config.Endpoint)

	return config, nil
}

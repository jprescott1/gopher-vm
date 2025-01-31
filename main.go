package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func main() {
	log.Print("Aint got no job")

	// Read yaml file
	yamlData, err := os.ReadFile("definition.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal yaml data
	var config Config
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Print data
	log.Println("Host:", config.Host)
}

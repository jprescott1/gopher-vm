package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Gopher VM")
}

func openFile() {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Println("Error reading YAML file: ", err)
		return
	}
	fmt.Println("YAML file: ", string(yamlFile))
}

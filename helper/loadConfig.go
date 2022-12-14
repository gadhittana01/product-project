package helper

import (
	"io/ioutil"
	"log"

	"gihub.com/gadhittana01/product-project/config"
	"gopkg.in/yaml.v2"
)

func LoadConfig(c *config.GlobalConfig) {
	path := "config/product-project.yaml"

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

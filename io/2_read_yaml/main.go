package main

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

type config struct {
	Id int64 `yaml:"id"`
	ProjectName string `yaml:"project_name"`
	Team struct {
		TeamName string `yaml:"team_name"`
		Membership int `yaml:"membership"`
	} `yaml:"team"`
	IssueBoard string `yaml:"issue_board"`
}

func (c *config) getConfig(fileName string) *config {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Error Reading yaml File: #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal Error: %v", err)
	}
	return c
}

func main() {
	var c config
	c.getConfig("config.yaml")

	fmt.Println(c)
}

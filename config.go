package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadConfig(
	Path string,
) *DatabaseConfig {
	// Read the config file
	file, err := os.Open(Path)
	if err != nil {
		panic(err)
	}
	// marshal the yaml file into the database config
	byts, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	db := &DatabaseConfig{}
	err = yaml.Unmarshal(byts, db)
	if err != nil {
		panic(err)
	}
	return db
}

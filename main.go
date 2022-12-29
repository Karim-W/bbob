package main

import "flag"

var (
	// Version is the version of the application
	Version = flag.Bool("version", false, "The version of the application")
	// Config is the path to the config file
	Config = flag.String("config", "config.yaml", "The path to the config file")
)

func main() {
	flag.Parse()
	if *Version {
		println("v0.0.1")
		return
	}
	// Read the config file
	conf := ReadConfig(*Config)
	// Setup the databases
	conf.Setup()
	// setup the logger
	// setup the database(s)
}

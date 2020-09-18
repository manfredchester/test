package main

import "flag"

// Command Command element
type Command struct {
	UUID    string
	File    string
	CldType string
	CldID   string
}

var cmd Command

// Init init command
func Init() {
	flag.StringVar(&cmd.UUID, "UUID", "", "accountID  such as 88")
	flag.StringVar(&cmd.File, "file", "", "filename  such as process.csv")
	flag.StringVar(&cmd.CldType, "type", "", "cloud type such as azure")
	flag.StringVar(&cmd.CldID, "cldID", "", "cloudid such as V570")
	flag.Parse()
}

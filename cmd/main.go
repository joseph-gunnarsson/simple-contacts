package main

import (
	"log"
	"os"

	command "github.com/joseph-gunnarsson/simple-contacts/cmd/contacts"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Error: No command provided. Please specify the configuration content as a command-line argument.")
	}

	switch os.Args[1] {
	case "add":
		command.AddContactCmd(os.Args[2:])
	case "update":
		command.UpdateContactCmd(os.Args[2:])
	case "delete":
		command.DeleteContactCmd(os.Args[2:])
	case "find":
		command.FindContactCmd(os.Args[2:])
	case "list":
		command.ListContactCmd()
	default:
		os.Exit(1)
	}
}

package command

import (
	"flag"
	"fmt"
	"log"
	"net/mail"
	"os"

	"github.com/joseph-gunnarsson/simple-contacts/internal/contacts"
)

func AddContactCmd(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	name := addCmd.String("name", "", "Contact name")
	phone := addCmd.String("phone", "", "Contact phone number")
	email := addCmd.String("email", "", "Contact email address")
	err := addCmd.Parse(args)
	if err != nil {
		log.Fatalf("Error parsing flags: %v\n", err)
	}

	if *name == "" || *phone == "" || *email == "" {
		addCmd.Usage()
		os.Exit(1)
	}

	if _, err := mail.ParseAddress(*email); err != nil {
		log.Fatalf("Invalid email address: %v\n", err)
	}
	err = contacts.AddContact(*name, *phone, *email)
	if err != nil {
		log.Fatalf("Error adding contact: %v\n", err)
	}
	fmt.Println("Contact added successfully")
}

func DeleteContactCmd(args []string) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	name := deleteCmd.String("name", "", "Contact name")
	err := deleteCmd.Parse(args)
	if err != nil {
		log.Fatalf("Error parsing flags: %v\n", err)
	}

	if *name == "" {
		deleteCmd.Usage()
		os.Exit(1)
	}

	err = contacts.DeleteContact(*name)
	if err != nil {
		log.Fatalf("Error adding contact: %v\n", err)
	}

	fmt.Println("Contact deleted successfully")
}

func FindContactCmd(args []string) {
	findCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	name := findCmd.String("name", "", "Contact name")
	err := findCmd.Parse(args)
	if err != nil {
		log.Fatalf("Error parsing flags: %v\n", err)
	}

	if *name == "" {
		findCmd.Usage()
		os.Exit(1)
	}

	contact, err := contacts.FindContact(*name)
	if err != nil {
		log.Fatalf("Error finding contact: %v\n", err)
	}

	fmt.Printf("Name: %s, Phone: %s, Email: %s\n", contact.Name, contact.PhoneNumber, contact.Email)
}

func UpdateContactCmd(args []string) {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	name := updateCmd.String("name", "", "Contact name")
	newName := updateCmd.String("newName", "", "Contact new name")
	phone := updateCmd.String("phone", "", "Contact phone number")
	email := updateCmd.String("email", "", "Contact email address")

	err := updateCmd.Parse(args)
	if err != nil {
		log.Fatalf("Error parsing flags: %v\n", err)
	}

	if *name == "" || (*phone == "" && *email == "" && *newName == "") {
		updateCmd.Usage()
		os.Exit(1)
	}

	if *phone != "" {
		err := contacts.UpdateContactPhone(*name, *phone)
		if err != nil {
			log.Fatalf("Error updating contact phone: %v\n", err)
		}
	}

	if *newName != "" {
		err := contacts.UpdateContactName(*name, *newName)
		if err != nil {
			log.Fatalf("Error updating contact name: %v\n", err)
		}
	}

	if *email != "" {
		if _, err := mail.ParseAddress(*email); err != nil {
			log.Fatalf("Invalid email address: %v\n", err)
		}

		err := contacts.UpdateContactEmail(*name, *email)
		if err != nil {
			log.Fatalf("Error updating contact email: %v\n", err)
		}
	}

	fmt.Println("Contact updated successfully")
}

func ListContactCmd() {
	contactsList := contacts.LoadContacts()

	if len(contactsList) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	for _, contact := range contactsList {
		fmt.Printf("Name: %s, Phone: %s, Email: %s\n", contact.Name, contact.PhoneNumber, contact.Email)
	}
}

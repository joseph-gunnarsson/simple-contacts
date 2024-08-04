package contacts

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func getContactFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".contact-data.json"), nil
}

func LoadContacts() []Contact {
	var contactArr []Contact

	contactData := getContactFile()
	if err := json.Unmarshal(contactData, &contactArr); err != nil {
		log.Fatalf("Error unmarshalling contacts: %v", err)
	}

	return contactArr
}

func getContactFile() []byte {
	contactFilePath, err := getContactFilePath()
	if err != nil {
		log.Fatalf("Error getting contact file path: %v", err)
	}

	contactData, err := os.ReadFile(contactFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			contactData = []byte("[]")
			createContactFile(contactFilePath, contactData)
		} else {
			log.Fatalf("Error reading contact file: %v", err)
		}
	}
	return contactData
}

func createContactFile(contactFilePath string, data []byte) {
	file, err := os.OpenFile(contactFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}

func SaveContacts(contacts []Contact) error {
	contactFilePath, err := getContactFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(contacts)
	if err != nil {
		return err
	}

	createContactFile(contactFilePath, data)

	return nil
}

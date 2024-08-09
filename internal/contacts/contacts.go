package contacts

import (
	"fmt"
)

type Contact struct {
	Name        string `json:"Name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

func AddContact(name, phone, email string) error {
	contactArr := LoadContacts()

	if _, exists := contactExists(contactArr, name); exists {
		return fmt.Errorf("contact with the name %s already exists", name)
	}

	newContact := Contact{
		Name:        name,
		PhoneNumber: phone,
		Email:       email,
	}

	contactArr = append(contactArr, newContact)

	err := SaveContacts(contactArr)
	if err != nil {
		return fmt.Errorf("failed to save contacts: %v", err)
	}

	return nil
}
func FindContact(name string) (Contact, error) {
	contactArr := LoadContacts()
	idx, exists := contactExists(contactArr, name)
	if exists {
		return Contact{}, fmt.Errorf("contact with the name %s does not exists", name)
	}

	return contactArr[idx], nil
}

func contactExists(contactArr []Contact, name string) (int, bool) {
	for i, contact := range contactArr {
		if contact.Name == name {
			return i, true
		}
	}
	return -1, false
}

func DeleteContact(name string) error {
	contactArr := LoadContacts()

	idx, exists := contactExists(contactArr, name)
	if !exists {
		return fmt.Errorf("contact with the name %s does not exist", name)
	}
	contactArr = append(contactArr[:idx], contactArr[idx+1:]...)
	SaveContacts(contactArr)
	return nil
}

func UpdateContactPhone(name, newPhone string) error {
	contactArr := LoadContacts()

	idx, exists := contactExists(contactArr, name)
	if !exists {
		return fmt.Errorf("contact with the name %s does not exist", name)
	}

	contactArr[idx].PhoneNumber = newPhone

	err := SaveContacts(contactArr)
	if err != nil {
		return fmt.Errorf("failed to save contacts: %v", err)
	}

	return nil
}

func UpdateContactEmail(name, newEmail string) error {
	contactArr := LoadContacts()

	idx, exists := contactExists(contactArr, name)
	if !exists {
		return fmt.Errorf("contact with the name %s does not exist", name)
	}

	contactArr[idx].Email = newEmail

	err := SaveContacts(contactArr)
	if err != nil {
		return fmt.Errorf("failed to save contacts: %v", err)
	}

	return nil
}

func UpdateContactName(oldName, newName string) error {
	contactArr := LoadContacts()

	idx, exists := contactExists(contactArr, oldName)
	if !exists {
		return fmt.Errorf("contact with the name %s does not exist", oldName)
	}

	contactArr[idx].Name = newName

	err := SaveContacts(contactArr)
	if err != nil {
		return fmt.Errorf("failed to save contacts: %v", err)
	}

	return nil
}

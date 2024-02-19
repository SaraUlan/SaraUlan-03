// services/contact/internal/domain/contact.go

package domain

import (
	"fmt"
	"strings"
)

type Contact struct {
	ID          int
	FirstName   string
	LastName    string
	MiddleName  string
	PhoneNumber string
}

func NewContact(id int, firstName, lastName, middleName, phoneNumber string) (*Contact, error) {

	if !isValidPhoneNumber(phoneNumber) {
		return nil, fmt.Errorf("invalid phone number: %s", phoneNumber)
	}

	contact := &Contact{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		MiddleName:  middleName,
		PhoneNumber: phoneNumber,
	}

	contact.setFullName()

	return contact, nil
}

func (c *Contact) FullName() string {
	return fmt.Sprintf("%s %s %s", c.FirstName, c.LastName, c.MiddleName)
}

func (c *Contact) setFullName() {
	c.FirstName = strings.TrimSpace(c.FirstName)
	c.LastName = strings.TrimSpace(c.LastName)
	c.MiddleName = strings.TrimSpace(c.MiddleName)
}

func isValidPhoneNumber(phoneNumber string) bool {
	for _, char := range phoneNumber {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

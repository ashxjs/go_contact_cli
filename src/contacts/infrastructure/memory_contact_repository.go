package infrastructure

import (
	"contact_cli/src/contacts/domain"
	"errors"
	"strings"
)

type MemoryContactRepository struct {
	contacts []domain.Contact
}

func NewMemoryContactRepository() *MemoryContactRepository {
	return &MemoryContactRepository{
		contacts: []domain.Contact{
			{ID: "1", FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Phone: "1234567890"},
			{ID: "2", FirstName: "Jane", LastName: "Smith", Email: "jane.smith@example.com", Phone: "0987654321"},
			{ID: "3", FirstName: "Alice", LastName: "Johnson", Email: "alice.johnson@example.com", Phone: "1122334455"},
		},
	}
}

func (r *MemoryContactRepository) GetContacts() ([]domain.Contact, error) {
	return r.contacts, nil
}

func (r *MemoryContactRepository) GetContactByFirstNameAndLastName(firstName string, lastName string) (domain.Contact, error) {
	for _, c := range r.contacts {
		if strings.EqualFold(c.FirstName, firstName) && strings.EqualFold(c.LastName, lastName) {
			return c, nil
		}
	}
	return domain.Contact{}, errors.New("contact not found")
}

func (r *MemoryContactRepository) CreateContact(contact domain.Contact) (domain.Contact, error) {
	r.contacts = append(r.contacts, contact)
	return contact, nil
}

package adapter

import (
	"contact_cli/src/contacts/application"
	"contact_cli/src/contacts/domain"
	"fmt"
)

type ContactHandler struct {
	service *application.ContactService
}

func NewContactHandler(service *application.ContactService) *ContactHandler {
	return &ContactHandler{service: service}
}

func (h *ContactHandler) GetContacts() {
	fmt.Println("GetContacts")
}

func (h *ContactHandler) GetContactByFirstNameAndLastName(firstName string, lastName string) (domain.Contact, error) {
	return h.service.GetContactByFirstNameAndLastName(firstName, lastName)
}

func (h *ContactHandler) CreateContact(contact domain.Contact) (domain.Contact, error) {
	return h.service.CreateContact(contact)
}

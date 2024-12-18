package application

import (
	"contact_cli/src/contacts/domain"
)

type ContactService struct {
	repo domain.ContactRepository
}

func NewContactService(repo domain.ContactRepository) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) GetContacts() ([]domain.Contact, error) {
	return s.repo.GetContacts()
}

func (s *ContactService) GetContactByFirstNameAndLastName(firstName string, lastName string) (domain.Contact, error) {
	return s.repo.GetContactByFirstNameAndLastName(firstName, lastName)
}

func (s *ContactService) CreateContact(contact domain.Contact) (domain.Contact, error) {
	return s.repo.CreateContact(contact)
}

package domain

type ContactRepository interface {
	GetContacts() ([]Contact, error)
	GetContactByFirstNameAndLastName(firstName string, lastName string) (Contact, error)
	CreateContact(contact Contact) (Contact, error)
}

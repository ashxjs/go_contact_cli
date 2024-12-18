package ui

import (
	"contact_cli/src/contacts/adapter"
	"contact_cli/src/contacts/domain"
	"fmt"
)

type ConsoleUI struct {
	contactHandler *adapter.ContactHandler
}

func NewConsoleUI(handler *adapter.ContactHandler) *ConsoleUI {
	return &ConsoleUI{
		contactHandler: handler,
	}
}

func (ui *ConsoleUI) Run() {
	for {
		command := ui.readCommand()
		if command == "exit" {
			break
		}
		ui.executeCommand(command)
	}
}

func (ui *ConsoleUI) readCommand() string {
	fmt.Println("Saisissez une commande:")
	var command string
	fmt.Scanln(&command)
	return command
}

func (ui *ConsoleUI) executeCommand(command string) {
	switch command {
	case "help":
		ui.printHelp()
	case "get":
		ui.handleGetContact()
	case "create":
		ui.handleCreateContact()
	default:
		fmt.Println("Unknown command")
	}
}

func (ui *ConsoleUI) handleGetContact() {
	firstName, lastName := ui.readFirstNameAndLastName()
	contact, err := ui.contactHandler.GetContactByFirstNameAndLastName(firstName, lastName)
	if err != nil {
		fmt.Println("Contact not found")
		return
	}
	contact.ToString()
}

func (ui *ConsoleUI) handleCreateContact() {
	firstName, lastName, email, phone := ui.readContactDetails()
	contact := domain.Contact{FirstName: firstName, LastName: lastName, Email: email, Phone: phone}
	createdContact, err := ui.contactHandler.CreateContact(contact)
	if err != nil {
		fmt.Println("Contact not created")
		return
	}
	createdContact.ToString()
}

// Helper methods moved from Program
func (ui *ConsoleUI) readContactDetails() (string, string, string, string) {
	firstName, lastName := ui.readFirstNameAndLastName()

	fmt.Println("Saisissez l'email du contact:")
	var email string
	fmt.Scanln(&email)

	fmt.Println("Saisissez le numéro de téléphone du contact:")
	var phone string
	fmt.Scanln(&phone)

	return firstName, lastName, email, phone
}

func (ui *ConsoleUI) readFirstNameAndLastName() (string, string) {
	fmt.Println("Saisissez le prénom du contact:")
	var firstName string
	fmt.Scanln(&firstName)

	fmt.Println("Saisissez le nom du contact:")
	var lastName string
	fmt.Scanln(&lastName)

	return firstName, lastName
}

func (ui *ConsoleUI) printHelp() {
	fmt.Println("\n********************")
	fmt.Println("ashx contact cli")
	fmt.Println("\n********************")
	fmt.Println("get: Récupérer un contact par son prénom et son nom")
	fmt.Println("create: Créer un contact")
	fmt.Println("exit: Quitter le programme")
	fmt.Println("help: Afficher l'aide")
	fmt.Println("\n********************")
}

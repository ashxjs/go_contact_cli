package domain

import (
	"contact_cli/src/contacts/adapter"
	"contact_cli/src/contacts/application"
	"contact_cli/src/contacts/domain"
	"contact_cli/src/contacts/infrastructure"
	ui "contact_cli/src/ui/domain"
)

// Program represents the core application orchestrator
type Program struct {
	contactRepository domain.ContactRepository
	contactService    *application.ContactService
	contactHandler    *adapter.ContactHandler
}

func NewProgram() *Program {
	contactRepository := infrastructure.NewMemoryContactRepository()
	contactService := application.NewContactService(contactRepository)
	contactHandler := adapter.NewContactHandler(contactService)

	return &Program{
		contactHandler:    contactHandler,
		contactRepository: contactRepository,
		contactService:    contactService,
	}
}

// Start initializes and runs the application
func (p *Program) Start() {
	ui := ui.NewConsoleUI(p.contactHandler)
	ui.Run()
}

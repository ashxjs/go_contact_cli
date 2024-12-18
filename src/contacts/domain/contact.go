package domain

import (
	"fmt"
	"time"
)

type Contact struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Contact) ToString() string {
	fmt.Println("\n********************")
	fmt.Println("FirstName: ", c.FirstName)
	fmt.Println("LastName: ", c.LastName)
	fmt.Println("Email: ", c.Email)
	fmt.Println("Phone: ", c.Phone)
	fmt.Println("********************")
	return ""
}

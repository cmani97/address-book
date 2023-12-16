package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	addbook "github.com/cmani97/address-book/src/application/addressbook"
)

func RunCLI() {
	ab := addbook.NewAddressBook()

	for {
		fmt.Println("1. Add Contact")
		fmt.Println("2. Search Contacts")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addContact(ab)
		case "2":
			searchContacts(ab)
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addContact(ab *addbook.AddressBook) {
	fmt.Println("Enter contact details:")
	// fmt.Print("First Name: ")
	// firstName := getUserInput()
	// fmt.Print("Last Name: ")
	// lastName := getUserInput()
	// fmt.Print("Address: ")
	// address := getUserInput()
	// fmt.Print("Phone Number: ")
	// phoneNumber := getUserInput()

	contact := &addbook.Contact{}
	fmt.Print("Enter First Name: ")
	contact.FirstName = getUserInput()

	fmt.Print("Enter Last Name: ")
	contact.LastName = getUserInput()

	fmt.Print("Enter Address: ")
	contact.Address = getUserInput()

	fmt.Print("Enter Phone Number: ")
	contact.PhoneNumber = getUserInput()

	ab.AddContact(contact)

	fmt.Println("Contact added successfully!")
}

func searchContacts(ab *addbook.AddressBook) {
	fmt.Print("Enter search term: ")
	searchTerm := getUserInput()

	fmt.Println("Search Results:")

	isPhoneNumber := strings.ContainsAny(searchTerm, "0123456789")
	fmt.Println("isPhoneNumber:", isPhoneNumber)
	if isPhoneNumber {
		matchingContactsByPhoneNumber := ab.SearchContactsByPhoneNumber(searchTerm)
		err := printContacts(matchingContactsByPhoneNumber, "Phone Number")
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		matchingContactsByName := ab.SearchContactsByName(searchTerm)
		err := printContacts(matchingContactsByName, "Name")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func printContacts(contacts []*addbook.Contact, searchType string) error {
	if len(contacts) == 0 {
		return fmt.Errorf("No contacts found for the given %s.\n", searchType)
	}

	fmt.Printf("Contacts found for %s:\n", searchType)
	for _, contact := range contacts {
		fmt.Printf("Name: %s %s, Address: %s, Phone Number: %s\n", contact.FirstName, contact.LastName, contact.Address, contact.PhoneNumber)
	}

	return nil
}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

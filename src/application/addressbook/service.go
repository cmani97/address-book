package addressbook

import (
	"sync"
)

type Contact struct {
	FirstName   string
	LastName    string
	Address     string
	PhoneNumber string
}

type TrieNode struct {
	children map[rune]*TrieNode
	contacts []*Contact
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
	}
}

type AddressBook struct {
	nameTrie     *TrieNode
	phoneHashMap map[string]*Contact
	lock         sync.RWMutex
}

func NewAddressBook() *AddressBook {
	return &AddressBook{
		nameTrie:     newTrieNode(),
		phoneHashMap: make(map[string]*Contact),
	}
}

func (ab *AddressBook) AddContact(contact *Contact) {
	ab.lock.Lock()
	defer ab.lock.Unlock()

	current := ab.nameTrie
	for _, char := range contact.FirstName + " " + contact.LastName {
		if _, exists := current.children[char]; !exists {
			current.children[char] = newTrieNode()
		}
		current = current.children[char]
	}
	current.contacts = append(current.contacts, contact)

	ab.phoneHashMap[contact.PhoneNumber] = contact
}

func (ab *AddressBook) SearchContactsByName(name string) []*Contact {
	ab.lock.RLock()
	defer ab.lock.RUnlock()

	current := ab.nameTrie
	for _, char := range name {
		if _, exists := current.children[char]; !exists {
			return []*Contact{}
		}

		current = current.children[char]
	}

	return current.contacts
}

func (ab *AddressBook) SearchContactsByPhoneNumber(phoneNumber string) []*Contact {
	ab.lock.RLock()
	defer ab.lock.RUnlock()

	if contact, ok := ab.phoneHashMap[phoneNumber]; ok {
		return []*Contact{contact}
	}

	return []*Contact{}
}

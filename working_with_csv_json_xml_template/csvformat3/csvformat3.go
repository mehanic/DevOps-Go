package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

// Initialize the logger
func init() {
	logFile := "/tmp/contacts.log"
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Repository manages CSV file operations for contacts
type Repository struct {
	fullFilePath string
}

// NewRepository creates a new Repository instance
func NewRepository(fullFilePath string) *Repository {
	log.Printf("Initializing contacts repository with file at: %s", fullFilePath)
	return &Repository{fullFilePath: fullFilePath}
}

// Add adds a new contact to the CSV file
func (r *Repository) Add(contact map[string]string) {
	log.Printf("Creating contact: %v", contact)

	headers := make([]string, 0, len(contact))
	for h := range contact {
		headers = append(headers, h)
	}
	sort.Strings(headers)

	writeHeaders := !fileExists(r.fullFilePath)

	file, err := os.OpenFile(r.fullFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if writeHeaders {
		log.Println("This is the first contact in the given file. Writing headers.")
		if err := writer.Write(headers); err != nil {
			log.Fatalf("Failed to write headers: %v", err)
		}
	}

	row := make([]string, len(headers))
	for i, h := range headers {
		row[i] = contact[h]
	}

	if err := writer.Write(row); err != nil {
		log.Fatalf("Failed to write row: %v", err)
	}
}

// Names retrieves all contact names from the CSV file
func (r *Repository) Names() []string {
	log.Println("Retrieving all contact names")

	file, err := os.Open(r.fullFilePath)
	if err != nil {
		log.Printf("Failed to open file: %v", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		return nil
	}

	if len(records) < 2 {
		return nil // No data or only headers
	}

	names := []string{}
	for _, record := range records[1:] {
		names = append(names, record[0]) // assuming the first column is "name"
	}

	log.Printf("Found %d contacts", len(names))
	return names
}

// FullContact returns the full contact details for a given name
func (r *Repository) FullContact(name string) map[string]string {
	log.Printf("Retrieving full contact for name: %s", name)

	file, err := os.Open(r.fullFilePath)
	if err != nil {
		log.Printf("Failed to open file: %v", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		return nil
	}

	headers := records[0]
	for _, record := range records[1:] {
		if record[0] == name {
			log.Println("Contact was found")
			contact := make(map[string]string)
			for i, header := range headers {
				contact[header] = record[i]
			}
			return contact
		}
	}

	log.Printf("Contact was not found for name: %s", name)
	return nil
}

// Main represents the main menu and logic
type Main struct {
	repo *Repository
}

// NewMain creates a new instance of Main
func NewMain(contactsFile string) *Main {
	return &Main{repo: NewRepository(contactsFile)}
}

// Create prompts the user for input and adds a new contact
func (m *Main) Create() {
	var name, number string
	fmt.Print("name: ")
	fmt.Scan(&name)
	fmt.Print("number: ")
	fmt.Scan(&number)
	contact := map[string]string{"name": name, "number": number}
	m.repo.Add(contact)
	fmt.Println("----------------")
}

// Names lists all the contact names
func (m *Main) Names() {
	names := m.repo.Names()
	if len(names) > 0 {
		for _, n := range names {
			fmt.Println("- " + n)
		}
	} else {
		fmt.Println("no contacts were found")
	}
	fmt.Println("----------------")
}

// FullContact prints the full contact details for a given name
func (m *Main) FullContact() {
	var name string
	fmt.Print("name: ")
	fmt.Scan(&name)
	contact := m.repo.FullContact(name)
	if contact != nil {
		fmt.Printf("name: %s\n", contact["name"])
		fmt.Printf("number: %s\n", contact["number"])
	} else {
		fmt.Println("contact not found.")
	}
	fmt.Println("----------------")
}

// Menu displays the main menu
func (m *Main) Menu() {
	for {
		fmt.Println("1) Create Contact")
		fmt.Println("2) All Contacts")
		fmt.Println("3) Detail of a contact")
		fmt.Println("0) Exit")

		var choice string
		fmt.Print("What do you want to do? ")
		fmt.Scan(&choice)

		switch choice {
		case "1":
			m.Create()
		case "2":
			m.Names()
		case "3":
			m.FullContact()
		case "0":
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

// Helper function to check if a file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !os.IsNotExist(err)
}

func main() {
	contactsFile := filepath.Join("/tmp", "contacts.csv")
	m := NewMain(contactsFile)
	m.Menu()
}

// 1) Create Contact
// 2) All Contacts
// 3) Detail of a contact
// 0) Exit
// What do you want to do? 2
// no contacts were found
// ----------------
// 1) Create Contact
// 2) All Contacts
// 3) Detail of a contact
// 0) Exit
// What do you want to do? 1
// name: max
// number: 3456
// ----------------
// 1) Create Contact
// 2) All Contacts
// 3) Detail of a contact
// 0) Exit
// What do you want to do? 2
// - max
// ----------------
// 1) Create Contact
// 2) All Contacts
// 3) Detail of a contact
// 0) Exit
// What do you want to do? 3
// name: max
// name: max
// number: 3456
// ----------------
// 1) Create Contact
// 2) All Contacts
// 3) Detail of a contact
// 0) Exit
// What do you want to do? 0

package main

import "fmt"
import "strings"
import "encoding/json"

type Employee struct {
	FirstName string
	LastName  string `json:"last_name"`
	Age       int
}

// Methode qui peut profiter de la structure
func (emp Employee) Email() string {
	return strings.ToLower(emp.FirstName + "." + emp.LastName + "@zenika.com")
}

//Basics
func main() {
	employees := []Employee{
		{
			FirstName: "Alain",
			LastName:  "Pipin",
			Age:       23,
		},
		{
			FirstName: "Vivi",
			LastName:  "Milou",
			Age:       26,
		},
		{
			FirstName: "Vivi",
			LastName:  "Milou",	
		},
	}

	for _, emp := range employees {
		email := emp.Email()
		// si on met lastName avec une minuscule, disparait du json
		fmt.Printf("Hello, %s %s %d ans\n", emp.FirstName, emp.LastName, emp.Age)
		fmt.Printf("Your email is, %s\n", email)
		// ! le package et la variable ont le même nom
		json, err := json.Marshal(emp)
		if err == nil {
			fmt.Println(string(json))
		}
	}
}

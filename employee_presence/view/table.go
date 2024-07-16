package view

import (
	"employee_presence/model"
	"fmt"
	"strings"
)

func Table(data []model.Employee) {

	fmt.Printf("%-5s %-20s %-10s %-20s\n", "ID", "Name", "Presence", "Created At")
	fmt.Println(strings.Repeat("-", 55))

	for _, emp := range data {
		fmt.Printf("%-5d %-20s %-10t %-20s\n", emp.ID, emp.Name, emp.Presence, emp.CreatedAt)
	}
}

func Menu() {
	fmt.Println("[1] View all employee presence")
	fmt.Println("[2] Create new employee")
	fmt.Println("[3] Delete employee by id")
	fmt.Println("[4] Update employee by id")
	fmt.Println("[5] Exit the program")
}

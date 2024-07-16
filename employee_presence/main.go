package main

import (
	"bufio"
	"employee_presence/model"
	"employee_presence/repository"
	"employee_presence/view"
	"fmt"
	"os"
)

func main() {
	// Init
	pr := repository.NewPresenceRepository()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\n")
	view.Menu()

	for {
		fmt.Print("\n")
		fmt.Print("Menu: ")
		input := view.InputCli(scanner)
		fmt.Print("\n")

		switch input {
		case "1":
			data := pr.GetAll()
			view.Table(data)
		case "2":
			fmt.Println("--Add new employee presence--")
			fmt.Print("Name: ")
			name := view.InputCli(scanner)

			employee := model.Employee{Name: name}

			pr.Save(employee)
			fmt.Println("Success saved: ", name)

		case "3":
			fmt.Println("--Delete employee presence by id--")
			fmt.Print("ID: ")
			id, err := view.InputCliNumber(scanner)

			if err != nil {
				fmt.Println("Error: should input a number")
				continue
			}

			err = pr.Delete(*id)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				continue
			}
			fmt.Println("Success delete employee with id: ", *id)

		case "4":
			fmt.Println("--Update employee presence by id--")
			fmt.Print("ID: ")
			id, err := view.InputCliNumber(scanner)

			if err != nil {
				fmt.Println("Error: should input a number")
				continue
			}
			employee, err := pr.FindById(*id)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				continue
			}

			fmt.Println()
			view.Table([]model.Employee{*employee})
			fmt.Println()

			if employee.Presence {
				fmt.Print("Are you want update the presence to false? (y/n)")
			} else {
				fmt.Print("Are you want update the presence to true? (y/n)")
			}

			fmt.Println()
			fmt.Print("input: ")
			input = view.InputCli(scanner)

			if input == "y" || input == "Y" {

				employee.Presence = !employee.Presence
				err = pr.Update(*employee)

				if err != nil {
					fmt.Println("Error: ", err.Error())
					continue
				}

				fmt.Println("Success update employee with id: ", *id)

			} else if input == "n" || input == "N" {
				fmt.Println("Cancel update employee id: ", id)
			} else {
				fmt.Println("Your input wrong!")
			}

		case "5":
			fmt.Println("Thankyou!")
			return
		default:
			fmt.Println("your input wrong!")
		}

	}
}

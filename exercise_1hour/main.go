package main

import (
	"bufio"
	"fmt"
	"os"
)

const lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Select Menu")
	fmt.Println("[1] Soal 1 :: score check")
	fmt.Println("[2] Soal 2 :: total price with discount")
	fmt.Println("[3] Soal 3 :: vocal char check from lorem")
	fmt.Println("[4] Soal 4 :: factorial 30")
	fmt.Println("[5] Soal 5 :: find odd or even from a vocal number")
	fmt.Println("[6] Soal 6 :: exit")
	for {
		fmt.Print("Select Menu > ")
		input := InputCli(scanner)

		switch input {

		// 1
		case "1":
			fmt.Print("Input number: ")
			number, err := InputCliNumber(scanner)
			if err != nil {
				fmt.Println("Input must a number!")
			}

			result := checkScore(number)

			fmt.Println("Result: ", result)

			// 2
		case "2":
			fmt.Print("Input Price: ")
			price, err := InputCliNumber(scanner)
			if err != nil {
				fmt.Println("Input must a number!")
			}

			fmt.Print("Input Quantity: ")
			quantity, err := InputCliNumber(scanner)
			if err != nil {
				fmt.Println("Input must a number!")
			}

			result := calculateTotal(*price, *quantity)
			fmt.Println("Result total price: ", result)

		//3
		case "3":
			fmt.Println("Calculate total vocal from this lorem")
			fmt.Println(lorem)
			result := calculateVocalChar(lorem)

			fmt.Println("Total vocal char: ", result)

		case "4":
			fmt.Println("Calculate factorial 30")
			result := factorial(30)

			fmt.Println("Result: ", result)

		case "5":
			fmt.Println("Find odd and even a number")
			fmt.Print("Input number in vocal: ")
			numberStr := InputCli(scanner)

			number := convertStrToInt(numberStr)

			if number == 0 {
				fmt.Println("Please input a correct number in vocal, or must less than 11")
				continue
			}

			oddOrEven := findOddOrEven(number)

			fmt.Printf("Your number %s is %s ", numberStr, oddOrEven)
			fmt.Println()
		}
	}
}

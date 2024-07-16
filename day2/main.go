package main

import (
	"bufio"
	"day2/model"
	"day2/repository"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	pr := repository.NewProductRepository()

	product1 := model.Product{
		ID:    1,
		Name:  "Indomie Goreng",
		Stock: 2,
	}

	product2 := model.Product{
		ID:    2,
		Name:  "Indomie",
		Stock: 5,
	}

	err := pr.Save(product1)

	err = pr.Save(product2)

	if err != nil {
		log.Println("Error save :: ", err.Error())
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Hello! Welcome from inventory app")
	fmt.Println("[1] Input 1 to get all products")
	fmt.Println("[2] Input 2 to insert new product")
	fmt.Println("[3] Input 3 to exit the program")

	for {

		fmt.Print("Input: ")
		scanner.Scan()

		input := scanner.Text()
		input = strings.TrimSpace(input)

		if input == "1" {
			fmt.Println("ID \t", "Name \t\t\t", "Stock \t")

			for _, product := range pr.FindAll() {

				fmt.Print(product.ID)
				fmt.Print("\t")

				fmt.Print(product.Name)
				fmt.Print("\t")
				fmt.Print("\t")

				fmt.Print(product.Stock)
				fmt.Print("\t")

				fmt.Println()
			}

		} else if input == "2" {
			var inputProduct model.Product
			fmt.Println("Input new product")
			fmt.Print("Insert name: ")

			scanner.Scan()
			inputProduct.Name = scanner.Text()

			fmt.Print("Insert stock: ")
			scanner.Scan()
			stock := strings.TrimSpace(scanner.Text())

			inputProduct.Stock, err = strconv.Atoi(stock)

			if err != nil {
				fmt.Println("Error input wrong :: ", stock)
			}

			lastProduct := pr.GetLastProduct()

			inputProduct.ID = lastProduct.ID + 1

			err := pr.Save(inputProduct)
			if err != nil {
				fmt.Println("Error save product :: ", err)
			}

			fmt.Println("Success saved product :: ", inputProduct)

		} else if input == "3" {
			fmt.Println("Thankyou! see u again")
			return
		} else {
			fmt.Println("Your input wrong!!!")

			fmt.Println("[1] Input 1 to get all products")
			fmt.Println("[2] Input 2 to insert new product")
			fmt.Println("[3] Input 3 to exit the app")
		}
	}
}

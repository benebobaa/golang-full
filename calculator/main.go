package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := NewResult()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Calculator")

	var state []string

	fmt.Println("[1] input 1 to Addition")
	fmt.Println("[2] input 2 to Subtraction")
	fmt.Println("[3] input 3 to Division")
	fmt.Println("[4] input 4 to Multiplication")
	fmt.Println("[5] input 5 to delete state")
	fmt.Println("[6] input 6 to exit the program")

	for {
		for _, v := range state {
			fmt.Print(v)
		}
		fmt.Println()
		if len(state) > 1 {
			fmt.Println("Result: ", r.Value)
			fmt.Println("Result Decimal: ", r.ValueDecimal)
		}
		fmt.Print("Select menu: ")
		scanner.Scan()

		input := scanner.Text()
		input_trim := strings.TrimSpace(input)

		if input_trim == "1" {
			var a, b string
			var bFloat float64

			fmt.Print("input number 1: ")
			scanner.Scan()
			a = scanner.Text()

			if len(state) < 2 {
				fmt.Print("input number 2: ")
				scanner.Scan()
				b = scanner.Text()
			}

			if len(state) >= 3 {
				state = append(state, " + ")
			}

			state = append(state, a)

			if len(state) <= 3 {
				state = append(state, " + ")
				state = append(state, b)
			}

			aFloat, err := strconv.ParseFloat(a, 64)

			if len(state) <= 3 {
				bFloat, err = strconv.ParseFloat(b, 64)
			} else {
				bFloat = r.ValueDecimal
			}

			if err != nil {
				fmt.Println("Your input must be number!")
				continue
			}

			r.Addition(aFloat, bFloat)

		} else if input_trim == "2" {
			var a, b string
			var bFloat float64

			fmt.Print("input number 1: ")
			scanner.Scan()
			a = scanner.Text()

			if len(state) < 2 {
				fmt.Print("input number 2: ")
				scanner.Scan()
				b = scanner.Text()
			}

			if len(state) >= 3 {
				state = append(state, " - ")
			}

			state = append(state, a)

			if len(state) <= 3 {
				state = append(state, " - ")
				state = append(state, b)
			}

			aFloat, err := strconv.ParseFloat(a, 64)

			if len(state) <= 3 {
				bFloat, err = strconv.ParseFloat(b, 64)
			} else {
				bFloat = r.ValueDecimal
			}

			if err != nil {
				fmt.Println("Your input must be number!")
				continue
			}

			r.Subtraction(aFloat, bFloat)

		} else if input_trim == "3" {
			var a, b string
			var bFloat float64

			fmt.Print("input number 1: ")
			scanner.Scan()
			a = scanner.Text()

			if len(state) < 2 {
				fmt.Print("input number 2: ")
				scanner.Scan()
				b = scanner.Text()
			}

			if len(state) >= 3 {
				state = append(state, " / ")
			}

			state = append(state, a)

			if len(state) <= 3 {
				state = append(state, " / ")
				state = append(state, b)
			}

			aFloat, err := strconv.ParseFloat(a, 64)

			if len(state) <= 3 {
				bFloat, err = strconv.ParseFloat(b, 64)
			} else {
				bFloat = r.ValueDecimal
			}

			if err != nil {
				fmt.Println("Your input must be number!")
				continue
			}

			r.Division(aFloat, bFloat)

		} else if input_trim == "4" {

		} else if input_trim == "5" {
			fmt.Println("Deleting state ...")
			state = []string{}
		} else if input_trim == "6" {
			fmt.Println("See u!")
			return
		} else {

			fmt.Println("your input wrong! please try again")
		}

	}

}

type Result struct {
	Value        int
	ValueDecimal float64
}

// type Operator interface {
// 	Addition(a, b float32) float32
// 	Subtraction(a, b float32) float32
// 	Division(a, b float32) (float32, error)
// 	Multiplication(a, b float32) float32
// }

func NewResult() *Result {
	return &Result{}
}

func (r *Result) Addition(a, b float64) {
	result := a + b
	r.Value = int(result)
	r.ValueDecimal = result
}

func (r *Result) Subtraction(a, b float64) {
	result := b - a
	r.Value = int(result)
	r.ValueDecimal = result
}

func (r *Result) Division(a, b float64) {
	result := a / b
	r.Value = int(result)
	r.ValueDecimal = result
}

//
// func (r *Result) Multiplication(a, b float32) float32 {
// 	result := a + b
// 	r.Value = int(result)
// 	r.ValueDecimal = result
// 	return result
// }

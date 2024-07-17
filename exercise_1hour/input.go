package main

import (
	"bufio"
	"strconv"
	"strings"
)

func InputCli(scanner *bufio.Scanner) string {

	scanner.Scan()

	input := scanner.Text()

	input = strings.TrimSpace(input)

	return input
}

func InputCliNumber(scanner *bufio.Scanner) (*int, error) {

	scanner.Scan()

	input := scanner.Text()

	input = strings.TrimSpace(input)

	number, err := strconv.Atoi(input)

	if err != nil {
		return nil, err
	}

	return &number, nil
}

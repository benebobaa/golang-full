package main

func factorial(number uint64) uint64 {

	if number == 0 || number == 1 {
		return 1
	}

	return number * factorial(number-1)
}

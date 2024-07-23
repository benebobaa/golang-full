package main

func convertStrToInt(str string) int {
	switch str {
	case "satu":
		return 1
	case "dua":
		return 2
	case "tiga":
		return 3
	case "empat":
		return 4
	case "lima":
		return 5
	case "enam":
		return 6
	case "tujuh":
		return 7
	case "delapan":
		return 8
	case "sembilan":
		return 9
	case "sepuluh":
		return 10
	default:
		return 0
	}
}

func findOddOrEven(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

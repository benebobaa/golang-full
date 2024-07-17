package main

func checkScore(number *int) string {
	if *number >= 90 {
		return "Selamat! Anda mendapatkan nilai A"
	} else if *number >= 80 && *number <= 89 {
		return "Anda mendapatkan nilai B"
	} else if *number >= 70 && *number <= 79 {
		return "Anda mendapatkan nilai C"
	} else if *number >= 60 && *number <= 69 {
		return "Anda mendapatkan nilai D"
	} else if *number < 60 {
		return "Anda mendapatkan nilai E"
	}
	return "Number out of range"
}

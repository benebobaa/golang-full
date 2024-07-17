package main

func calculateTotal(price, quantity int) float32 {
	// 0.9 mean 90%, because we want it get total after
	// calculate discount 10%
	return float32(price) * float32(quantity) * 0.9
}

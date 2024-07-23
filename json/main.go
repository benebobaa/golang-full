package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type TestAja struct {
	A, B string
	C    *string
}

func main() {

	r, err := json.Marshal(TestAja{
		A: "a",
		B: "b",
	})

	fmt.Println("bytes: ", r)
	fmt.Println("r: ", string(r))
	fmt.Println("err: ", err)

	for i, v := range r {
		if v == 97 {
			r[i] = 99
			break
		}
	}

	a, _ := os.Open("")

	json.NewDecoder(a)

	var test TestAja
	err = json.Unmarshal(r, &test)

	fmt.Println("test: ", test)
}

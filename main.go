package main

import "fmt"

func main() {
	numeros := make([]int, 100)
	for i := 0; i < 100; i++ {
		numeros[i] = i + 1
	}

	for _, num := range numeros {
		if num % 3 == 0 {
			fmt.Println(num)
		}
	}
}
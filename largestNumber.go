package main

import "fmt"

func largest(args ...int) (largest int) {
	largest = args[0]
	for _, num := range args {
		if num > largest {
			largest = num
		}
	}
	return
}

func main() {
	fmt.Println(largest(1, 20, 4, 7, 10))
}

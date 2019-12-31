/*
	A function that generates oddNumbers
*/

package main

import "fmt"

func generateOddNumber() func() int {
	oddNum := 1

	return func() (ret int) {
		ret = oddNum
		oddNum += 2
		return
	}
}

func main() {
	makeOdd := generateOddNumber()
	makeOdd()
	fmt.Println(makeOdd())
	fmt.Println(makeOdd())
}

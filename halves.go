/* function that takes an integer and halves it and returns true if it was even
or false if it was odd.
*/

package main

import "fmt"

func half(num int) (float64, bool) {
	halfOfNum := float64(num) / 2
	even := num%2 == 0
	return halfOfNum, even
}

func main() {

	var num int
	fmt.Scanf("%d", &num)
	fmt.Println(half(num))
}

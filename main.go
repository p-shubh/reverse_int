package main

import "fmt"

func reverse(num int) int {

	// var result int
	temp := 0
	if num > 0 {
		for i := num; i > 0; i = i / 10 {
			remainder := i % 10
			temp = (temp * 10) + remainder

		}

	} else {
		for i := num; i < 0; i = i / 10 {
			remainder := i % 10
			temp = (temp * 10) + remainder

		}
		// result = temp
	}
	return temp
}
func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}

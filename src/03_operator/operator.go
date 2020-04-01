package main

import "fmt"

func main() {

	var num1 uint8 = 3
	var num2 uint8 = 2

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)
	fmt.Println()
	fmt.Println(num1 << num2)
	fmt.Println(num1 >> num2)
	fmt.Println(^num1)

	var num3 int = 3
	var num4 float32 = 2.2

	fmt.Println()
	// fmt.Println(num3 + num4) comfile error
	fmt.Println(num3 + int(num4))



}
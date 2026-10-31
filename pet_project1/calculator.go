package main

import "fmt"

func main() {
	var num1, num2 float32
	var action rune
	fmt.Scanf("%f %f %c", &num1, &num2, &action)
	switch action {
	case '+':
		fmt.Println(num1 + num2)
	case '-':
		fmt.Println(num1 - num2)
	case '*':
		fmt.Println(num1 * num2)
	case '/':
		if num2 == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Println(num1 / num2)
		}
	case '%':
		fmt.Println(int(num1) % int(num2))
	}
}

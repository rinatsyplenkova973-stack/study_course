package main

import "fmt"

func main() {
	numbers()
}
func hello() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if (b - a) == (c - b) {
		fmt.Println("Да")
	} else {
		fmt.Println("Нет")
	}
}

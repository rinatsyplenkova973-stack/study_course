package main

import "fmt"

func main() {
	var age int
	fmt.Scan(&age)
	var result bool
	result = age >= 13 && age <= 19
	fmt.Println(result)
}

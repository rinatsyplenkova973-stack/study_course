package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fact := 1
	for i := 1; i <= a; i++ {
		fact = fact * i
	}
	fmt.Println(fact)
}

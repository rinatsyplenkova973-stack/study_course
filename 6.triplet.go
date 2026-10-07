package main

import "fmt"

func main() {
	var a, b, c rune
	fmt.Scan(&a, &b, &c)
	fmt.Println(string(a) + string(b) + string(c))
}

package main

import "fmt"

func sum() {
	var a int
	sum := 0
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}

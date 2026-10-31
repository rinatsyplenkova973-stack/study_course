package main

import "fmt"

func mult() {
	var a int
	fmt.Scan(&a)
	for i := 1; i <= 10; i++ {
		fmt.Println(a, "x", i, "=", a*i)
	}
}

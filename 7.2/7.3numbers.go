package main

import "fmt"

func numbers() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if a > b && a > c && a > d {
		fmt.Println(a)
	} else if b > a && b > c && b > d {
		fmt.Print(b)
	} else if c > a && c > b && c > d {
		fmt.Print(c)
	} else if d > a && d > c && d > b {
		fmt.Print(d)
	}

}

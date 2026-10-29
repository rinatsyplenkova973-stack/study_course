package main

import "fmt"

func main() {
	var attempts int
	fmt.Scan(&attempts)
	if attempts%2 == 0 {
		fmt.Println("любит")
	} else {
		fmt.Println("не любит")
	}
}

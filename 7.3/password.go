package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var text string
	fmt.Scan(&text)
	var length int = utf8.RuneCountInString(text)
	if length <= 8 {
		fmt.Println("Легкий")
	} else if length > 8 && length <= 12 {
		fmt.Println("Средний")
	} else if length > 12 && length <= 16 {
		fmt.Println("Сложный")
	} else if length > 16 {
		fmt.Println("Очень сложный")
	}
}

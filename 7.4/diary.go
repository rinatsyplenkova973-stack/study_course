package main

import "fmt"

func main() {
	calender()
	var note int
	fmt.Scan(&note)
	switch note {
	case 1:
		fmt.Println("Очень плохо")
	case 2:
		fmt.Println("Плохо")
	case 3:
		fmt.Println("Удовлетворительно")
	case 4:
		fmt.Println("Хорошо")
	case 5:
		fmt.Println("Отлично")
	default:
		fmt.Println("Некорректная оценка")
	}

}

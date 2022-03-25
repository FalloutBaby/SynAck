package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("Выберите количество потоков: ")
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Збазиба!")
}

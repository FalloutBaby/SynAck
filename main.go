package main

import (
	"fmt"
	"net/url"
)

func main() {
	var countGo int
	var address, port string

	fmt.Print("Выберите количество потоков: ")
	_, err := fmt.Scanf("%d", &countGo)

	if err != nil {
		fmt.Println("Gorutines is invalid,", err)
		return
	}

	fmt.Print("Введите адрес прозвона: ")
	_, err = fmt.Scanf("%s", &address)
	if err != nil {
		fmt.Println("Adress is Invalid", err)
		return
	}

	_, err = url.ParseRequestURI(address)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Введите порт прозвона: ")
	_, err = fmt.Scanf("%s", &port)

	if err != nil {
		fmt.Println("Port is invalid,", err)
		return
	}

	fmt.Println("Збазиба!")
}

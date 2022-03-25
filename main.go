package main

import (
	"SynAck/requests"
	"fmt"
	"net/url"
)

func main() {
	var grt int
	var addr, p string

	fmt.Print("Выберите количество потоков: ")
	_, err := fmt.Scanf("%d", &grt)

	if err != nil {
		fmt.Println("Goroutines is invalid,", err)
		return
	}

	fmt.Print("Введите адрес прозвона: ")
	_, err = fmt.Scanf("%s", &addr)
	if err != nil {
		fmt.Println("Address is Invalid", err)
		return
	}

	_, err = url.Parse(addr)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Введите порт прозвона: ")
	_, err = fmt.Scanf("%s", &p)

	if err != nil {
		fmt.Println("Port is invalid,", err)
		return
	}

	requests.Call(addr, p, grt)

	fmt.Println("Збазиба!")
}

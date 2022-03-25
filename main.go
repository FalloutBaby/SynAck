package main

import (
	"SynAck/worker"
	"fmt"
	"net/url"
)

func main() {
	var grt int
	var addr, p string

	fmt.Print("Выберите количество потоков: ")
	_, err := fmt.Scanln(&grt)

	if err != nil {
		fmt.Println("Goroutines is invalid,", err)
		return
	}

	fmt.Print("Введите адрес прозвона: ")
	_, err = fmt.Scanln(&addr)
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

	worker.Scan(addr, p, grt)

	fmt.Println("Збазиба!")
}

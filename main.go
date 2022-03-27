package main

import (
	"SynAck/producer"
	"SynAck/worker"
	"fmt"
	"net/url"
)

func main() {
	var grt int
	var addr string

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

	p := producer.GetPorts()

	scan := worker.Worker{}.Scan(addr, p, grt)

	fmt.Println(scan)

	fmt.Println("Збазиба!")
}

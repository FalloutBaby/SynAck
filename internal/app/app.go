package app

import (
	"SynAck/internal/services/decorators"
	"SynAck/internal/services/producers"
	"SynAck/internal/services/workers"
	"fmt"
	"regexp"
)

func Run() {
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
	var isValid bool
	isValid, err = regexp.MatchString("([/:\\w\\d]{2,256}\\.)+[\\w]{2,4}", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !isValid {
		panic("Url is invalid")
	}

	p := producers.GetPorts()

	dialer := decorators.NetDialer{}
	decorator := decorators.NetDecorator{Dialer: dialer}
	openPs := workers.Worker{Decorator: decorator}.Scan(addr, p, grt)

	fmt.Println(openPs)
	fmt.Println("Збазиба!")
}

package delivery

import (
	"fmt"
	"regexp"
)

type Delivery interface {
	GetAddress() string
	GetNetwork() string
}

type Http struct {
}

func (http Http) GetAddress() string {
	var addr string
	fmt.Print("Введите адрес прозвона: ")
	_, err := fmt.Scanln(&addr)
	if err != nil {
		panic(err)
	}
	var isValid bool
	isValid, err = regexp.MatchString("([/:\\w\\d]{2,256}\\.)+[\\w]{2,4}", addr)
	if err != nil {
		panic(err)
	}
	if !isValid {
		panic("Url is invalid")
	}
	return addr
}

const tcp = "tcp"

func (http Http) GetNetwork() string {
	return tcp
}

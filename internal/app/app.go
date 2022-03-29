package app

import (
	"SynAck/internal/delivery"
	"SynAck/internal/services/decorators"
	"SynAck/internal/services/producers"
	"SynAck/internal/services/workers"
	"fmt"
)

type App struct {
	dialer    decorators.NetDialer
	decorator decorators.NetDecorator
	producer  producers.Generator
	delivery  delivery.Http
}

func Run() {
	app := new(App)
	decorator := decorators.NetDecorator{Dialer: app.dialer}

	worker := workers.Worker{Decorator: decorator, Delivery: app.delivery, Producer: &app.producer}
	openPs := worker.ScanPorts()

	fmt.Println(openPs)
	fmt.Println("Збазиба!")
}

package main

import (
	"SynAck/internal/app"
	"os"
)

func main() {
	app.Run(os.Args[1], os.Args[2])
}

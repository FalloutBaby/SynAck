package decorators

import (
	"fmt"
	"testing"
)

func TestDial(t *testing.T) {
	network := "tcp"
	address := "scanme.nmap.org"

	d := new(Decorator)
	conn, err := d.Dial(network, address)
	if err != nil {
	}
	fmt.Println(conn)
}

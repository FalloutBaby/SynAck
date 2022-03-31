package decorators

import (
	"net"
	"strconv"
	"time"
)

type Dialer interface {
	DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)
}

type NetDialer struct {
}

func (d NetDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout(network, address, timeout) // TODO: внутри есть структура Dialer, которую можно было мочить без декоратора
}

type DialerDecorator interface { // TODO: Это не совсем декоратор это адаптер или фасад
	DialPort(network, addr string, p int) int
}

type NetDecorator struct {
	Dialer Dialer
}

func (d NetDecorator) DialPort(network, addr string, p int) int {
	timeout := time.Second
	c, err := d.Dialer.DialTimeout(network, addr+":"+strconv.Itoa(p), timeout)
	if err != nil {
		return 0
	} else { // TODO: else лишний, только увеличивает вложеность
		err = c.Close()
		if err != nil {
			panic(err)
		}
		return p
	}
}

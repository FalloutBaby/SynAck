package decorators

import (
	"net"
	"time"
)

type Dialer interface {
	DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)
}

type NetDialer struct {
}

func (d NetDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout(network, address, timeout)
}

type DialerDecorator interface {
	DialAll(network, addr string, p int) string
}

type NetDecorator struct {
	Dialer Dialer
}

func (d NetDecorator) DialAll(network, addr string, p int) string {
	timeout := time.Second
	c, err := d.Dialer.DialTimeout(network, addr+":"+string(p), timeout)
	if err != nil {
		//TODO: change
		return ""
	} else {
		err := c.Close()
		if err != nil {
			panic(err)
		}
		return string(p)
	}
}

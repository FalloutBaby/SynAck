package decorators

import (
	"net"
	"strings"
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
	DialAll(network, addr string, ps []string) string
}

type NetDecorator struct {
	Dialer Dialer
}

func (d NetDecorator) DialAll(network, addr string, ps []string) string {
	var result []string
	timeout := time.Second
	for _, p := range ps {
		c, err := d.Dialer.DialTimeout(network, addr+":"+p, timeout)
		if err != nil {
			continue
		} else {
			err := c.Close()
			if err != nil {
				panic(err)
			}
			result = append(result, p)
		}
	}
	return strings.Join(result, ", ")
}

package decorators

import (
	"log"
	"net"
	"strings"
)

type Dialer interface {
	Dial(network, address string) (net.Conn, error)
}

type NetDialer struct {
}

func (d NetDialer) Dial(network, address string) (net.Conn, error) {
	return net.Dial(network, address)
}

type DialerDecorator interface {
	DialAll(network, addr string, ps []string) string
}

type NetDecorator struct {
	Dialer Dialer
}

func (d NetDecorator) DialAll(network, addr string, ps []string) string {
	var result []string
	for _, p := range ps {
		c, err := d.Dialer.Dial(network, addr+":"+p)
		if err != nil {
			continue
		} else {
			err := c.Close()
			if err != nil {
				log.Panic(err)
			}
			result = append(result, p)
		}
	}
	return strings.Join(result, ", ")
}
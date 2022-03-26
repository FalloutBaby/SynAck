package decorators

import "net"

type Decorator struct {
	net.Dialer
	network, address string
}

func (d Decorator) Dial(network, address string) (net.Conn, error) {
	conn, err := d.Dial(network, address)
	if err != nil {

	}
	return conn, nil
}

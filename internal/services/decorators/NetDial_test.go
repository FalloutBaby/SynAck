package decorators

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
)

type ConnStub struct {
	isClosed bool
}

func (c ConnStub) Read(b []byte) (n int, err error) {
	return 0, nil
}
func (c ConnStub) Write(b []byte) (n int, err error) {
	return 0, nil
}
func (c ConnStub) Close() error {
	if c.isClosed {
		return errors.New("failed close connection")
	}
	return nil
}
func (c ConnStub) LocalAddr() net.Addr {
	return nil
}
func (c ConnStub) RemoteAddr() net.Addr {
	return nil
}
func (c ConnStub) SetDeadline(t time.Time) error {
	return nil
}
func (c ConnStub) SetReadDeadline(t time.Time) error {
	return nil
}
func (c ConnStub) SetWriteDeadline(t time.Time) error {
	return nil
}

type TestDialer struct {
	isOpen bool
	conn   ConnStub
}

func (d TestDialer) Dial(network, address string) (net.Conn, error) {
	if !d.isOpen {
		return d.conn, errors.New("failed connection")
	}
	return d.conn, nil
}

type dialAllDataProvider struct {
	dialer    TestDialer
	network   string
	address   string
	ports     []string
	openPorts string
}

func TestDialAll(t *testing.T) {
	provider := []dialAllDataProvider{
		{
			TestDialer{isOpen: true, conn: ConnStub{isClosed: false}},
			"tcp",
			"scanme.nmap.org",
			[]string{"80", "280", "443", "488", "591", "593", "623", "664", "777", "832", "1128"},
			"80, 280, 443, 488, 591, 593, 623, 664, 777, 832, 1128",
		},
		{
			TestDialer{isOpen: false, conn: ConnStub{isClosed: false}},
			"tcp",
			"scanme.nmap.org",
			[]string{"80", "280", "443", "488", "591", "593", "623", "664", "777", "832", "1128"},
			"",
		},
	}
	for _, p := range provider {
		ps := NetDecorator{p.dialer}.DialAll(p.network, p.address, p.ports)
		assert.Equal(t, p.openPorts, ps)
	}
}

type PanicDataProvider struct {
	dialer  TestDialer
	network string
	address string
	ports   []string
}

func TestDialAllWhenPanic(t *testing.T) {
	p := PanicDataProvider{
		TestDialer{isOpen: true, conn: ConnStub{isClosed: true}},
		"tcp",
		"scanme.nmap.org",
		[]string{"80", "280", "443", "488", "591", "593", "623", "664", "777", "832", "1128"},
	}
	assert.Panics(t, func() {
		NetDecorator{p.dialer}.DialAll(p.network, p.address, p.ports)
	})
}

package worker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type DialerStub struct{}

func (DialerStub) DialAll(network, addr string, ps []string) string {
	for _, e := range ps {
		if e == "80" {
			return "80"
		}
	}
	return ""
}

type dataProvider struct {
	network     string
	addr        string
	grt         int
	ports       []string
	resultPorts []string
}

func TestScan(t *testing.T) {
	tests := dataProvider{
		addr: "scanme.nmap.org",
		grt:  5,
		ports: []string{"80", "280", "443", "488", "591", "593", "623",
			"664", "777", "832", "1128", "1129", "1183", "1184",
			"5000", "5001", "8008", "8080", "11371"},
		resultPorts: []string{"80"},
	}
	w := Worker{dialer: DialerStub{}}
	result := w.Scan(tests.addr, tests.ports, tests.grt)
	assert.Equal(t, tests.resultPorts, result)
}

package workers

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type DialerStub struct {
	openPs    []string
	callCount int
}

func (d *DialerStub) DialAll(network, addr string, ps []string) string {
	d.callCount += 1
	currentOpenPs := strings.Join(ps, ", ")
	d.openPs = append(d.openPs, currentOpenPs)
	return currentOpenPs
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
	dialer := DialerStub{}
	w := Worker{Decorator: &dialer}
	result := w.Scan(tests.addr, tests.ports, tests.grt)
	assert.Equal(t, dialer.openPs, result)
	assert.Equal(t, dialer.callCount, tests.grt)
}

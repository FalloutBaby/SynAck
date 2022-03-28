package workers

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type DialerStub struct {
	network string
	addr    string
	openPs  []string
	cntGs   int
}

func (d *DialerStub) DialAll(network, addr string, ps []string) string {
	d.network = network
	d.addr = addr
	d.cntGs += 1
	openPs := strings.Join(ps, ", ")
	d.openPs[d.cntGs-1] = openPs

	return openPs
}

type dataProvider struct {
	network string
	addr    string
	grt     int
	ps      []string
}

func TestScan(t *testing.T) {
	tests := dataProvider{
		network: "tcp",
		addr:    "scanme.nmap.org",
		grt:     5,
		ps: []string{"80", "280", "443", "488", "591", "593", "623",
			"664", "777", "832", "1128", "1129", "1183", "1184",
			"5000", "5001", "8008", "8080", "11371"},
	}

	dialer := DialerStub{openPs: make([]string, tests.grt)}
	w := Worker{Decorator: &dialer}
	result := w.Scan(tests.addr, tests.ps, tests.grt)

	assert.Equal(t, dialer.network, tests.network)
	assert.Equal(t, dialer.addr, tests.addr)
	assert.Equal(t, dialer.cntGs, tests.grt)
	for _, exp := range dialer.openPs {
		assert.Contains(t, result, exp)
	}
}

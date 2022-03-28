package workers

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type DialerStub struct {
	openPs []string
}

type DeliveryStub struct {
}

type ProducerStub struct {
}

func (ps ProducerStub) GetGorutines() int {
	return 5
}

func (ds DeliveryStub) GetAddress() string {
	return "scanme.nmap.org"
}

func (ds DeliveryStub) GetNetwork() string {
	return "tcp"
}

func (d *DialerStub) DialAll(network, addr string, ps []string) string {
	currentOpenPs := strings.Join(ps, ", ")
	d.openPs = append(d.openPs, currentOpenPs)

	return currentOpenPs
}

func TestScan(t *testing.T) {
	ps := []string{"80", "280", "443", "488", "591", "593", "623",
		"664", "777", "832", "1128", "1129", "1183", "1184",
		"5000", "5001", "8008", "8080", "11371"}

	dialer := &DialerStub{}
	delivery := DeliveryStub{}
	producer := ProducerStub{}
	w := Worker{Decorator: dialer, Delivery: delivery, Producer: producer}

	result := w.Scan(ps)

	for _, exp := range dialer.openPs {
		assert.Contains(t, result, exp)
	}
}

package workers

import (
	"SynAck/internal/services/producers"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DialerStub struct {
	openPs []int
}

type DeliveryStub struct {
}

type ProducerStub struct {
	psChan chan int
	producers.Generator
}

func (ps *ProducerStub) WritePsToChan(psChan *chan int) {
	for i := 1; i <= cap(*psChan); i++ {
		*psChan <- i
	}
}

func (ds DeliveryStub) GetAddress() string {
	return "scanme.nmap.org"
}

func (ds DeliveryStub) GetNetwork() string {
	return "tcp"
}

func (d *DialerStub) DialPort(network, addr string, p int) int {
	d.openPs = append(d.openPs, p)
	return p
}

func TestScanPorts(t *testing.T) {
	i := 25
	psChan := make(chan int, i)

	dialer := &DialerStub{}
	delivery := DeliveryStub{}
	producer := ProducerStub{}
	w := Worker{Decorator: dialer, Delivery: delivery, Producer: &producer}

	producer.WritePsToChan(&psChan)

	result := w.ScanPorts("", 5)

	for _, exp := range dialer.openPs {
		assert.Contains(t, result, exp)
	}
}

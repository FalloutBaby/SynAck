package workers

import (
	"SynAck/internal/services/producers"
	"fmt"
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

func (ps ProducerStub) GetGorutines() int {
	return 5
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

func (d *DialerStub) DialAll(network, addr string, p int) string {
	d.openPs = append(d.openPs, p)

	return fmt.Sprint(p)
}

func TestScanPorts(t *testing.T) {
	psChan := make(chan int, 25)

	dialer := &DialerStub{}
	delivery := DeliveryStub{}
	producer := ProducerStub{}
	w := Worker{Decorator: dialer, Delivery: delivery, Producer: &producer}

	producer.WritePsToChan(&psChan)

	result := w.ScanPorts()

	for _, exp := range dialer.openPs {
		assert.Contains(t, result, exp)
	}
}

package workers

import (
	"SynAck/internal/services/decorators"
	"math"
	"sync"
)

const tcpNetwork = "tcp"

type Worker struct {
	Decorator decorators.DialerDecorator
}

func (w Worker) Scan(addr string, ports []string, grt int) []string {
	wg := sync.WaitGroup{}

	splitPs := splitPorts(ports, grt)

	chanel := make(chan string, grt)
	for i := 0; i < grt; i++ {
		ps := splitPs[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			dial := w.Decorator.DialAll(tcpNetwork, addr, ps)
			chanel <- dial
		}()
	}

	var result []string
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < grt; i++ {
			ps := <-chanel
			if ps != "" {
				result = append(result, ps)
			}
		}
	}()
	wg.Wait()

	return result
}

func splitPorts(ports []string, grt int) [][]string {
	countPorts := int(math.Ceil(float64(len(ports)) / float64(grt)))

	var splitPs [][]string

	for i := 0; i < len(ports); i += countPorts {
		end := i + countPorts

		if end >= len(ports) {
			end = len(ports) - 1
		}

		splitPs = append(splitPs, ports[i:end])
	}
	return splitPs
}

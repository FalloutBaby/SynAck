package worker

import (
	"SynAck/decorators"
	"math"
	"sync"
)

const tcpNetwork = "tcp"

type Worker struct {
	dialer decorators.DialerDecorator
}

func (w Worker) Scan(addr string, ports []string, grt int) []string {
	wg := sync.WaitGroup{}

	splitPs := splitPorts(ports, grt)

	var result []string
	for i := 0; i < grt; i++ {
		ps := splitPs[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			dial := w.dialer.DialAll(tcpNetwork, addr, ps)
			if dial != "" {
				result = append(result, dial)
			}
		}()
	}
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

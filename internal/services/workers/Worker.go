package workers

import (
	"SynAck/internal/delivery"
	"SynAck/internal/services/decorators"
	"SynAck/internal/services/producers"
	"sync"
)

type Worker struct {
	Decorator decorators.DialerDecorator
	Delivery  delivery.Delivery
	Producer  producers.Producer
}

func (w Worker) ScanPorts(addr string, grt int) []int {
	tcp := w.Delivery.GetNetwork()
	amount := 81
	psChan := make(chan int, amount)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		w.Producer.WritePsToChan(&psChan)
	}()

	chanel := make(chan int, cap(psChan)-1)
	for i := 0; i < grt; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			var cycle int
			var exit bool
			for {
				p, open := <-psChan
				if len(psChan) != 0 && open {
					dial := w.Decorator.DialPort(tcp, addr, p)
					chanel <- dial
				} else if len(psChan) == 0 && open {
					close(psChan)
					exit = true
					break
				} else {
					exit = true
					break
				}
				cycle++
			}

			if exit {
				return
			}
		}()
	}

	var result []int
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < cap(chanel); i++ {
			ps := <-chanel
			if ps != 0 {
				result = append(result, ps)
			}
		}
		close(chanel)
	}()
	wg.Wait()

	return result
}

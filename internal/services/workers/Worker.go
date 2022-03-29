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

func (w Worker) ScanPorts() []string {
	addr := w.Delivery.GetAddress()
	tcp := w.Delivery.GetNetwork()
	grt := w.Producer.GetGorutines()

	psChan := make(chan int, 80)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		w.Producer.WritePsToChan(&psChan)
	}()

	chanel := make(chan string, cap(psChan))
	for i := 0; i < grt; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if p, open := <-psChan; open {
					dial := w.Decorator.DialAll(tcp, addr, p)
					chanel <- dial
				} else if len(psChan) == 0 {
					close(psChan)
					break
				} else {
					break
				}
			}
		}()
	}

	var result []string
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= cap(chanel); i++ {
			ps := <-chanel
			if ps != "" {
				result = append(result, ps)
			}
		}
		close(chanel)
	}()
	wg.Wait()

	return result
}

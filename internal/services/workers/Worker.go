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
	tcp := w.Delivery.GetTcpNetwork()
	wg := sync.WaitGroup{}

	psChan := make(chan int, w.Producer.GetCountPorts()) // TODO: Зачем такой большой канал?
	go w.Producer.WritePsToChan(psChan)                  // TODO: Тут не нужен такой большой канал, так как нет цели записать сразу все, оттуда работу будут забирать
	var m sync.Mutex

	var result []int
	for i := 0; i < grt; i++ {
		wg.Add(1)
		go func() { // TODO: Почему не выделили в отдельную зависимость?
			defer wg.Done()
			for p := range psChan {
				dial := w.Decorator.DialPort(tcp, addr, p)
				if dial != 0 {
					m.Lock()
					result = append(result, dial)
					m.Unlock()
				}
			}
		}()
	}
	wg.Wait()

	return result
}

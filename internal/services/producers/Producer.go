package producers

type Producer interface {
	WritePsToChan(psChan chan int)
	GetCountPorts() int
}

type Generator struct {
}

func (g Generator) WritePsToChan(psChan chan int) {
	for i := 1; i <= cap(psChan); i++ {
		psChan <- i
	}
	close(psChan) // TODO: закрываться канал дожен там же где его открыли
}

func (g Generator) GetCountPorts() int {
	return 65536
}

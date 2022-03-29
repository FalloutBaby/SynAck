package producers

type Producer interface {
	WritePsToChan(psChan *chan int)
}

type Generator struct {
}

func (g *Generator) WritePsToChan(psChan *chan int) {
	for i := 1; i <= cap(*psChan); i++ {
		*psChan <- i
	}
}

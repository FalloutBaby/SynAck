package producers

import "fmt"

type Producer interface {
	GetGorutines() int
	WritePsToChan(psChan *chan int)
}

type Generator struct {
}

func (g *Generator) WritePsToChan(psChan *chan int) {
	psCnt := 65536
	for i := 1; i <= psCnt; i++ {
		*psChan <- i
	}
}

func (g Generator) GetGorutines() int {
	var grt int

	fmt.Print("Выберите количество потоков: ")
	_, err := fmt.Scanln(&grt)
	if err != nil {
		panic(err)
	}
	return grt
}

package worker

import (
	"fmt"
	"math"
	"net"
	"sync"
)

func Scan(addr string, ps []int, grt int) {
	wg := sync.WaitGroup{}
	//chs := make(map[int]chan string)

	psChunk := int(math.Ceil(float64(len(ps)) / float64(grt)))

	for i := 1; i <= grt; i++ {
		wg.Add(1)
		chs[i] = make(chan string)
		go func() {
			defer wg.Done()

			_, err := net.Dial("tcp", addr+":"+p)
			if err != nil {
				fmt.Println(err)
			}
			chs[i] <- p
			//TODO: expand
			println(p)
		}()
	}
	wg.Wait()
}

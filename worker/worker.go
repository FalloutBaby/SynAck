package worker

import (
	"fmt"
	"net"
	"sync"
)

func Scan(addr, p string, grt int) {
	wg := sync.WaitGroup{}
	chs := make(map[int]chan string)

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

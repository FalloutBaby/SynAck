package requests

import (
	"fmt"
	"net"
	"sync"
)

func Call(addr, p string, grt int) {
	wg := sync.WaitGroup{}
	wg.Add(grt)

	chs := make(map[int]chan string)
	for i := 1; i <= grt; i++ {
		chs[i] = make(chan string)
		go func() {
			defer wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println(chs)
	_, err := net.Dial("tcp", addr+":"+p)

	if err != nil {
		fmt.Println(err)
	}
}

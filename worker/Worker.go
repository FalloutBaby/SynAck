package worker

import (
	"fmt"
	"math"
	"net"
	"sync"
)

func Scan(addr string, ports []string, grt int) []string {
	wg := sync.WaitGroup{}

	countPorts := int(math.Ceil(float64(len(ports)) / float64(grt)))

	var chunkPs [][]string

	for i := 0; i < len(ports); i += countPorts {
		end := i + countPorts

		if end >= len(ports) {
			end = len(ports) - 1
		}

		chunkPs = append(chunkPs, ports[i:end])
	}

	var result []string
	for i := 1; i <= grt; i++ {
		ps := chunkPs[i-1]
		wg.Add(1)
		go func() {
			defer wg.Done()

			for _, p := range ps {
				_, err := net.Dial("tcp", addr+":"+p)
				if err != nil {
					fmt.Println(err)
				} else {
					result = append(result, p)
				}
			}
		}()
	}
	wg.Wait()

	return result
}

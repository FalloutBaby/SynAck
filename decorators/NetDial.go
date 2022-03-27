package decorators

import (
	"log"
	"net"
)

func Dial(network, addr string, ps []string) []string {
	var result []string
	for _, p := range ps {
		c, err := net.Dial(network, addr+":"+p)
		if err != nil {
			continue
		} else {
			err := c.Close()
			if err != nil {
				log.Panic(err)
			}
			result = append(result, p)
		}
	}
	return result
}

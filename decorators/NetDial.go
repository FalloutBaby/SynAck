package decorators

import "net"

func Dial(network, addr string, ps []string) []string {
	var result []string
	for _, p := range ps {
		_, err := net.Dial(network, addr+":"+p)
		if err != nil {
		} else {
			result = append(result, p)
		}
	}
	return result
}

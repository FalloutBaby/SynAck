package requests

import "net"

func Call(address, port string) {
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		println(err)
		return
	}
}

package requests

import "net"

func Call(port string) {
	conn, err := net.Dial("tcp", "scanme.nmap.org:"+port)
	if err != nil {
		println(err)
		return
	}
}

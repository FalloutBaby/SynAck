package decorators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDialWhenSuccess(t *testing.T) {
	network := "tcp"
	address := "scanme.nmap.org"
	ports := []string{"80", "280", "443", "488", "591", "593", "623", "664", "777", "832", "1128"}

	ps := Dial(network, address, ports)
	assert.Equal(t, []string{"80"}, ps)
}

func TestDial(t *testing.T) {
	network := "tcp"
	address := "scanme.nmap.org"
	ports := []string{"1129", "1183", "1184", "5000", "5001", "8008", "8080", "11371"}

	ps := Dial(network, address, ports)
	assert.Nil(t, ps)
}

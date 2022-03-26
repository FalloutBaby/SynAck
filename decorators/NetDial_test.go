package decorators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDial(t *testing.T) {
	network := "tcp"
	address := "scanme.nmap.org"

	ps := Dial(
		network,
		address,
		[]string{"80", "280", "443", "488", "591", "593", "623",
			"664", "777", "832", "1128", "1129", "1183", "1184",
			"5000", "5001", "8008", "8080", "11371"},
	)
	assert.Equal(t, ps, []string{"80"})
}

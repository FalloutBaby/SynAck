package producers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerator_WritePsToChan(t *testing.T) {
	psChan := make(chan int)
	g := Generator{}
	go g.WritePsToChan(&psChan)
	for i := 1; i <= 65536; i++ {
		assert.Equal(t, i, <-psChan)
	}
	close(psChan)
}

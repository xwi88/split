package split_test

import (
	"testing"

	"github.com/xwi88/split"
)

func Test_Split(t *testing.T) {
	data := "xwi88"
	count := 100
	mod := 100
	seed := 0x00

	for i := 0; i < count; i++ {
		bk := split.Split(data, mod, seed+i)
		t.Logf("data: %v, mod: %v, seed: %v, bucket:%d", data, mod, seed+i, bk)
	}
}

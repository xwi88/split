package split

import (
	"github.com/xwi88/split/murmur3"
)

var Split split

type split func(data string, mod, seed int) int

func init() {
	Split = murmur3.Split
}

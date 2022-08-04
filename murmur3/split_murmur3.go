package murmur3

import (
	mm3 "github.com/spaolacci/murmur3"
)

func Split(data string, mod, seed int) (b int) {
	mh := mm3.New32WithSeed(uint32(seed))
	_, _ = mh.Write([]byte(data))
	return int(mh.Sum32()%uint32(mod) + 1)
}

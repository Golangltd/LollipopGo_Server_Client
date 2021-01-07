package uint64

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetUint64() uint64 {
	return uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
}

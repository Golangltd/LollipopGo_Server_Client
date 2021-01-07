package temp

import "sync"

var cardUID uint64
var cardMutex sync.Mutex

func GernerateCardUID() uint64 {
	cardMutex.Lock()
	defer cardMutex.Unlock()
	cardUID++
	return cardUID
}

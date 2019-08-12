package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	m := map[int]struct{}{}
	var mu sync.RWMutex

	go func() {
		for {
			mu.Lock()
			m[rand.Intn(100000)] = struct{}{}
			mu.Unlock()
		}
	}()

	for {
		mu.RLock()
		v, ok := m[rand.Intn(100000)]
		fmt.Println(v, ok) // {} true と {} false が混ざって表示されるが、mから要素を消さないので次第に {} true ばっかり表示されるようになる
		mu.RUnlock()
	}
}

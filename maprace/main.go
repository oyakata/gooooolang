package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
)

func selectFromMySQL() map[int]struct{} {
	// データの全取っ替え。
	// ここではランダムにやっているけど、実際はバッチが定期的に更新するMySQLのテーブルをSELECTすると思ってください。
	size := rand.Intn(100)
	m := make(map[int]struct{}, size)
	for i := 0; i < size; i++ {
		m[i] = struct{}{}
	}
	return m
}

func main() {
	var at atomic.Value
	at.Store(selectFromMySQL())

	go func() {
		for {
			at.Store(selectFromMySQL())
		}
	}()

	for {
		m := at.Load().(map[int]struct{}) // Loadはinterface{}型を返すので型アサーションが必要
		v, ok := m[50]
		fmt.Println(v, ok) // {} true と {} false が混ざって表示される
	}
}

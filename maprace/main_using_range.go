package main

import (
	"fmt"
	"math/rand"
)

func main() {
	m := map[int]struct{}{}
	go func() {
		for {
			m[rand.Intn(100000)] = struct{}{}
		}
	}()

	for {
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}

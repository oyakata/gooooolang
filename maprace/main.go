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
		fmt.Println(m[rand.Intn(100000)])
	}
}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("===== START =====")

	m := map[int]struct{}{}
	go func() {
		for {
			m[rand.Intn(100000)] = struct{}{}
		}
	}()

	for {
		// fatal error: concurrent map read and map write
		fmt.Println(m[rand.Intn(100000)])

		// fatal error: concurrent map iteration and map write
		// for k, v := range m {
		//     fmt.Println(k, v)
		// }
	}
}

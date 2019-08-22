package main

import (
	"fmt"
)

type Foo struct {
	Name string
}

func (f *Foo) IsNull() bool {
	return f == nil
}

func (f *Foo) Hello() string {
	return "Hello, world."
}

func main() {
	var f *Foo
	fmt.Println(f.IsNull(), f.Hello(), f) // => true Hello, world. <nil>
}

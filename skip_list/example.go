package main

import (
    "fmt"
    "github.com/huandu/skiplist"
)

func main() {
	list := skiplist.New(skiplist.Int)

	list.Set(12, "hello world")
	list.Set(34, 36)

	elem := list.Get(34)
	fmt.Println(elem.Value)
	next := elem.Next()

	val, ok := list.GetValue(34)
	fmt.Println(val, ok)
	list.Remove(34)
}
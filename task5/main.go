package main

import "fmt"

type MyInt int

type Counter[T ~int | ~int8 | ~int16 | ~int32 | ~int64] struct {
	number T
}

func (c *Counter[T]) Next() {
	c.number++
}

// Task: create Counter with method Next(). Counter expected to support all kinds of int and its inherit types
func main() {
	c := Counter[MyInt]{
		number: MyInt(5),
	}
	c.Next()
	fmt.Println(c.number)
}

package main

import (
	"fmt"
)

type A struct{}

type B struct{}

func (A) String() string {
	return "a"
}

func (B) String() string {
	return "b"
}

func ToStringsA(s []A) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = v.String()
	}
	return r
}

func ToStringsB(s []B) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = v.String()
	}
	return r
}

func ToStrings(s []fmt.Stringer) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = v.String()
	}
	return r
}

func ToStringsI(s []fmt.Stringer) []string {
	r := make([]string, len(s))
	for i := range s {
		r[i] = "i"
	}
	return r
}

// Task: implement again functions ToStrings, ToStringsI for fmt.Stringers
func main() {
	arrA := []A{{}, {}}
	arrB := []B{{}, {}}

	fmt.Println(ToStringsA(arrA))
	fmt.Println(ToStringsB(arrB))

	arr := []fmt.Stringer{A{}, B{}}

	fmt.Println(ToStrings(arr))
	fmt.Println(ToStringsI(arr))
}

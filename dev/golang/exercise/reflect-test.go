package main

import (
	"fmt"
	"reflect"
)


func main() {

	var i int
	var ms map[string] int
	var mi map[string] interface {}
	var s string
	var sl [5]int
	var ss []string

	fmt.Println("int is ",reflect.TypeOf(i))
	fmt.Println("map of string-int is ",reflect.TypeOf(ms))
	fmt.Println("map of string-interface is ",reflect.TypeOf(mi))
	fmt.Println("string is ",reflect.TypeOf(s))
	fmt.Println("slice of int is ",reflect.TypeOf(sl))
	fmt.Println("slice of string is ",reflect.TypeOf(ss))



}

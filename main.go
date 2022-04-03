package main

import "reflect"

func main() {
	var v interface{} = true
	reflect.TypeOf(v)
}

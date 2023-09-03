package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("DSA Practice")
	x := []any{1, nil, 3}
	val := reflect.ValueOf(x[0])
	fmt.Println(val.IsValid())
}

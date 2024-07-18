package main

import "fmt"

type x interface{}

func main() {
	var x interface{} = 10
	var y interface{} = "Hello world"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel %T é o valor %v\n", t, t)
}

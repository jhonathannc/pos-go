package main

import "fmt"

var b bool
var (
	c int
	d string
	e float64
)

func main() {
	var meuarray [3]int
	meuarray[0]= 10
	meuarray[1]= 20
	meuarray[2]= 30

	for i,v := range meuarray {
		fmt.Printf("valor: %d, indice: %d\n", v, i)
	}
}
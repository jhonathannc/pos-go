package main

import (
	"fmt"

	"github.com/jhonathannc/packaging/3/math"
)

func main() {
	m := math.Math{A: 10, B: 20}
	fmt.Println(m.Sum())
}

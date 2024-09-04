package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jhonathannc/packaging/4/math"
)

func main() {
	m := math.Math{A: 10, B: 20}
	fmt.Println(m.Sum())

	println(uuid.New().String())
}

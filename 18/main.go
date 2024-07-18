package main

import (
	"fmt"

	"pos-go/matematica"

	"github.com/google/uuid"
)

func main() {
	res := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v", res)

	fmt.Println(uuid.New())
}

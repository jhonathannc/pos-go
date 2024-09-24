package main

import (
	"fmt"

	"github.com/jhonathannc/fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}

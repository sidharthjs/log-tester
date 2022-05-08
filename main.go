package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello")
	log.Info("App started")
}

func Adder(a, b int) int {
	return a + b
}

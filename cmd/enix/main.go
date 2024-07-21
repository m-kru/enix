package main

import (
	"github.com/m-kru/enix/internal/arg"
	"log"
)

func main() {
	log.SetFlags(0)

	arg.Parse()

	println("Hello enix")
}

package main

import (
	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/window"
	"log"
)

func main() {
	log.SetFlags(0)

	arg.Parse()

	colors, err := cfg.Init()
	if err != nil {
		log.Fatalf("%v", err)
	}

	window.Start()
}

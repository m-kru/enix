package main

import (
	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/enix"
	"log"
)

func main() {
	log.SetFlags(0)

	arg.Parse()

	config, colors, keys, err := cfg.Init()
	if err != nil {
		log.Fatalf("%v", err)
	}

	enix.Start(&config, &colors, &keys)
}

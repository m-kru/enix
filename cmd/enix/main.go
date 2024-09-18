package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/enix"
	"github.com/m-kru/enix/internal/script"
)

func main() {
	log.SetFlags(0)

	arg.Parse()

	config, colors, keys, promptKeys, insertKeys, err := cfg.Init()
	if err != nil {
		log.Fatalf("%v", err)
	}

	if arg.DumpConfig {
		data, err := json.MarshalIndent(config, "", "\t")
		if err != nil {
			log.Fatalf("%v", err)
		}
		fmt.Printf("%s\n", string(data))
		os.Exit(0)
	}

	if arg.Script != "" {
		err := script.Exec()
		if err != nil {
			log.Fatalf("%v", err)
		}
		os.Exit(0)
	}

	enix.Start(&config, &colors, &keys, &promptKeys, &insertKeys)
}

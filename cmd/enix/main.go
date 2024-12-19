package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/enix"
	"github.com/m-kru/enix/internal/script"
)

func main() {
	log.SetFlags(0)

	arg.Parse()

	keys, promptKeys, insertKeys, err := cfg.Init()
	if err != nil {
		log.Fatalf("%v", err)
	}

	if arg.DumpConfig {
		data, err := json.MarshalIndent(cfg.Cfg, "", "\t")
		if err != nil {
			log.Fatalf("%v", err)
		}
		fmt.Printf("%s\n", string(data))
	}

	if arg.DumpKeys {
		data, err := json.MarshalIndent(keys, "", "\t")
		if err != nil {
			log.Fatalf("%v", err)
		}
		fmt.Printf("%s\n", string(data))
	}

	if arg.DumpPromptKeys {
		data, err := json.MarshalIndent(promptKeys, "", "\t")
		if err != nil {
			log.Fatalf("%v", err)
		}
		fmt.Printf("%s\n", string(data))
	}

	if arg.DumpConfig || arg.DumpKeys || arg.DumpPromptKeys {
		os.Exit(0)
	}

	if arg.Profile {
		f, err := os.Create("enix.prof")
		if err != nil {
			log.Fatal("can't create cpu profile file: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("can't start cpu profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if arg.Script != "" {
		err := script.Exec()
		if err != nil {
			log.Fatalf("%v", err)
		}
		os.Exit(0)
	}

	enix.Start(&keys, &promptKeys, &insertKeys)
}

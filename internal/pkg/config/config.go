package config

import (
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

// ParseConfig parses all provided flags/envs and sets values into fields of an abstract config
func ParseConfig(config interface{}) interface{} {
	p := flags.NewParser(config, flags.Default)
	if _, err := p.Parse(); err != nil {
		log.Println("[ERROR] Error while parsing config options:", err)
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	return config
}

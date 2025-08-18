package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	yaml "github.com/goccy/go-yaml"
)

type MainFlags struct {
	ConfPath string
}

func LoadMainFlags(args []string) (*MainFlags, []string, error) {
	var mainFlags MainFlags
	fls := flag.NewFlagSet("", flag.ExitOnError)
	fls.StringVar(&mainFlags.ConfPath, "conf-path", "", "GLobal configuration path")
	err := fls.Parse(args)
	if err != nil {
		return nil, nil, fmt.Errorf("[LoadMainFlags] Unable to parse global falgs. %w", err)
	}
	return &mainFlags, fls.Args(), nil
}

func LoadConf(path string) (*Conf, error) {
	var conf Conf
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("[load-conf] Unable to load file %s. %w", path, err)
	}
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse file %s. %w", path, err)
	}
	return &conf, nil
}

func main() {
	conf, args, err := LoadMainFlags(os.Args[1:])
	if err != nil {
		log.Fatalf("[main] Unable to parse flags. %s", err)
	}
	if len(args) > 0 {
		switch args[0] {
		case "ls":
			LsCommand(conf, args[1:])
		case "run":
			RunCommand(conf, args[1:])
		}
	} else {
		log.Fatal("Missing command parameter")
	}
}

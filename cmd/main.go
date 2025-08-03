package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Service struct {
	Name string   `yaml:"name"`
	Exec string   `yaml:"exec"`
	Args []string `yaml:"args"`
}

type Conf struct {
	Services string `yaml:"services"`
}

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

func main() {
	nfls, args, err := LoadMainFlags(os.Args[1:])
	log.Printf("Conf path: %s", nfls.ConfPath)
	if err != nil {
		log.Fatalf("[main] Unable to parse flags. %s", err)
	}
	if len(args) > 1 {
		switch args[1] {
		case "ls":
			log.Println("ls command")
		case "run":
			log.Println("run command")
		}
	} else {
		log.Fatal("Missing command parameter")
	}
}

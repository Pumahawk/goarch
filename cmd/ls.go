package main

import (
	"flag"
	"fmt"
	"log"
)

type LsCommandFlags struct {
}

func LsCommand(mf *MainFlags, args []string) {
	_, args, err := LoadLsCommandFlags(args)
	if err != nil {
		log.Fatalf("[ls] Unable to parse flags. %s", err)
	}
	conf, err := LoadConf(mf.ConfPath)
	if err != nil {
		log.Fatalf("[ls] Unable to load conf file. %s", err)
	}
	for _, s := range conf.Services {
		fmt.Printf("%s\n", s.Name)
	}
}

func LoadLsCommandFlags(args []string) (*LsCommandFlags, []string, error) {
	var lsCommandFlags LsCommandFlags
	fls := flag.NewFlagSet("", flag.ExitOnError)
	err := fls.Parse(args)
	if err != nil {
		return nil, nil, fmt.Errorf("[LoadLsCommandFlags] Unable to parse falgs. %w", err)
	}
	return &lsCommandFlags, fls.Args(), nil
}

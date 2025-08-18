package main

import (
	"flag"
	"fmt"
	"log"
)

func RunCommand(mf *MainFlags, args []string) {
	_, args, err := LoadRunCommandFlags(args)
	if err != nil {
		log.Fatalf("[run] Unable to parse flags. %s", err)
	}
	conf, err := LoadConf(mf.ConfPath)
	if err != nil {
		log.Fatalf("[run] Unable to load conf file. %s", err)
	}
	if len(args) < 1 {
		log.Fatalf("[run] missin service name")
	}
	serviceNameRun := args[0]
	for _, s := range conf.Services {
		if s.Name == serviceNameRun {
			RunService(&s, args)
			return
		}
	}
	log.Fatalf("Service %s not found", serviceNameRun)
}

func LoadRunCommandFlags(args []string) (*LsCommandFlags, []string, error) {
	var lsCommandFlags LsCommandFlags
	fls := flag.NewFlagSet("", flag.ExitOnError)
	err := fls.Parse(args)
	if err != nil {
		return nil, nil, fmt.Errorf("[LoadRunCommandFlags] Unable to parse falgs. %w", err)
	}
	return &lsCommandFlags, fls.Args(), nil
}

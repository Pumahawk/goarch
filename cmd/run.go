package main

import (
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

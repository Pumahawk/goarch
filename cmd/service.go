package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

type Service struct {
	Name             string   `yaml:"name"`
	Exec             string   `yaml:"exec"`
	Args             []string `yaml:"args"`
	Tags             []string `yaml:"tags"`
	WorkingDirectory *string  `yaml:"working-directory"`
	StdIn            *string  `yaml:"stdIn"`
}

func RunService(service *Service, args []string) {
	allArgs := append(service.Args, args[1:]...)
	cmd := exec.Command(service.Exec, allArgs...)
	if service.WorkingDirectory != nil {
		cmd.Dir = *service.WorkingDirectory
	}
	var stdin io.Reader
	if service.StdIn == nil {
		stdin = os.Stdin
	} else {
		stdin = strings.NewReader(*service.StdIn)
	}
	cmd.Stdin = stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	csign := make(chan os.Signal, 1)
	signal.Notify(csign)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("[run] Unable to start process. %s", err)
	}
	go func() {
		for s := range csign {
			if cmd.Process != nil {
				cmd.Process.Signal(s)
			}
		}
	}()
	cmd.Wait()
	signal.Stop(csign)
	close(csign)
	os.Exit(cmd.ProcessState.ExitCode())
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


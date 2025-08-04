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

	yaml "github.com/goccy/go-yaml"
)

type Service struct {
	Name             string   `yaml:"name"`
	Exec             string   `yaml:"exec"`
	Args             []string `yaml:"args"`
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

type Conf struct {
	Services []Service `yaml:"services"`
}

type LsCommandFlags struct {
}

type MainFlags struct {
	ConfPath string
}

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

func LoadRunCommandFlags(args []string) (*LsCommandFlags, []string, error) {
	var lsCommandFlags LsCommandFlags
	fls := flag.NewFlagSet("", flag.ExitOnError)
	err := fls.Parse(args)
	if err != nil {
		return nil, nil, fmt.Errorf("[LoadRunCommandFlags] Unable to parse falgs. %w", err)
	}
	return &lsCommandFlags, fls.Args(), nil
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

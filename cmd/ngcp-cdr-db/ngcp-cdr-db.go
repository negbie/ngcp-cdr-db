package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/koding/multiconfig"
	"github.com/negbie/logp"
	"github.com/negbie/ngcp-cdr-db/config"
	"github.com/negbie/ngcp-cdr-db/scan"
)

const version = "0.81"

type processor interface {
	Run()
	End()
}

func init() {
	var err error
	var logging logp.Logging
	var fileRotator logp.FileRotator

	c := multiconfig.New()
	cfg := new(config.Configuration)
	c.MustLoad(cfg)
	config.Setting = *cfg

	if tomlExists(config.Setting.ConfigFile) {
		cf := multiconfig.NewWithPath(config.Setting.ConfigFile)
		err := cf.Load(cfg)
		if err == nil {
			config.Setting = *cfg
		} else {
			fmt.Printf("syntax error %v in ngcp-cdr-db.toml, use flag defaults\n", err)
		}
	} else {
		fmt.Println("couldn't find ngcp-cdr-db.toml, use flag defaults")
	}

	logp.DebugSelectorsStr = &config.Setting.LogDbg
	logging.Level = config.Setting.LogLvl
	logging.ToSyslog = &config.Setting.LogSys
	logp.ToStderr = &config.Setting.LogStd
	fileRotator.Path = "./"
	fileRotator.Name = "ngcp-cdr-db.log"
	logging.Files = &fileRotator

	err = logp.Init("ngcp-cdr-db", &logging)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func tomlExists(f string) bool {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	} else if !strings.Contains(f, ".toml") {
		return false
	}
	return err == nil
}

func main() {
	if config.Setting.Version {
		fmt.Println(version)
		os.Exit(0)
	}
	var wg sync.WaitGroup
	var sigCh = make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	csv := scan.NewCSV()
	processors := []processor{csv}

	for _, proc := range processors {
		wg.Add(1)
		go func(p processor) {
			defer wg.Done()
			p.Run()
		}(proc)
	}

	<-sigCh

	for _, proc := range processors {
		wg.Add(1)
		go func(p processor) {
			defer wg.Done()
			p.End()
		}(proc)
	}
	wg.Wait()
}

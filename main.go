package main

import (
	"flag"
	"fmt"
	"github.com/fantasticmao/csv-to-ical/app"
	"github.com/fantasticmao/csv-to-ical/config"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"
)

var (
	showVersion bool
	configDir   string
)

func init() {
	flag.BoolVar(&showVersion, "v", false, "show current version")
	flag.StringVar(&configDir, "d", "", "specify the configuration directory")
	flag.Parse()
}

func main() {
	if showVersion {
		fmt.Printf("%v %v %v-%v with %v at commit %v build %v\n", config.Name, config.Version,
			runtime.GOOS, runtime.GOARCH, runtime.Version(), config.CommitHash, config.BuildTime)
		return
	}

	if configDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fatal("get user home directory error: %v\n", err.Error())
		}
		configDir = path.Join(homeDir, ".config", config.Name)
	}
	appConfig, err := config.ParseConfig(path.Join(configDir, "config.yaml"))
	if err != nil {
		fatal("parse config file error: %v\n", err.Error())
	}

	for owner, provider := range appConfig.CsvProviders {
		err = app.RegisterHandler(owner, provider)
		fatal("register HTTP handler error: %v\n", err.Error())
	}

	app.StartServer(appConfig.BindAddress)
	fmt.Printf("start HTTP server success, bind address: %v\n", appConfig.BindAddress)

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}

func fatal(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

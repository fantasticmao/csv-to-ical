package main

import (
	"flag"
	"fmt"
	"github.com/fantasticmao/csv-to-ical/app"
	"github.com/fantasticmao/csv-to-ical/common"
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
		fmt.Printf("%v %v %v-%v with %v built on commit %v at %v\n", common.Name, common.Version,
			runtime.GOOS, runtime.GOARCH, runtime.Version(), common.CommitHash, common.BuildTime)
		return
	}

	if configDir == "" {
		if homeDir, err := os.UserHomeDir(); err != nil {
			fatal("get user home directory error: %v\n", err.Error())
		} else {
			configDir = path.Join(homeDir, ".config", common.Name)
		}
	}

	appConfig, err := common.ParseConfig(path.Join(configDir, "config.yaml"))
	if err != nil {
		fatal("parse config file error: %v\n", err.Error())
	}

	for owner, provider := range appConfig.CsvProviders {
		err = app.RegisterLocalHandler(configDir, owner, provider)
		if err != nil {
			fatal("register HTTP handler error: %v\n", err.Error())
		}
	}
	app.RegisterRemoteHandler()
	app.RegisterVersionHandler()
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

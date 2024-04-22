package main

import (
	"flag"
	"github.com/fantasticmao/csv-to-ical/app"
	"github.com/fantasticmao/csv-to-ical/common"
	"github.com/fantasticmao/csv-to-ical/log"
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
		log.Info("%v %v %v-%v with %v built on commit %v at %v", common.Name, common.Version,
			runtime.GOOS, runtime.GOARCH, runtime.Version(), common.CommitHash, common.BuildTime)
		return
	}

	if configDir == "" {
		if homeDir, err := os.UserHomeDir(); err != nil {
			log.Panic(err, "get user home directory error")
		} else {
			configDir = path.Join(homeDir, ".config", common.Name)
		}
	}

	appConfig, err := common.ParseConfig(path.Join(configDir, "config.yaml"))
	if err != nil {
		log.Panic(err, "parse config file error")
	}

	for owner, provider := range appConfig.CsvProviders {
		err = app.RegisterLocalHandler(configDir, owner, provider)
		if err != nil {
			log.Panic(err, "register HTTP handler error: %v")
		}
	}
	app.RegisterRemoteHandler()
	app.RegisterDefaultHandler()
	app.StartServer(appConfig.BindAddress)
	log.Info("start HTTP server success, bind address: %v", appConfig.BindAddress)

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}

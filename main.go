package main

import (
	"flag"
	"github.com/fantasticmao/csv-to-ical/app"
	"github.com/fantasticmao/csv-to-ical/common"
	"github.com/fantasticmao/csv-to-ical/log"
	"github.com/gin-gonic/gin"
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

	configPath := path.Join(configDir, "config.yaml")
	appConfig, err := common.ParseConfig(configPath)
	if err != nil {
		log.Panic(err, "parse config file error, config path: '%v'", configPath)
	}

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Content-Type", "text/plain; charset=UTF-8")
	})
	metrics := app.MetricMiddleware()

	r.GET("/", metrics, app.HomeHandler())
	r.GET("/version", metrics, app.VersionHandler())
	r.GET("/remote", metrics, app.RemoteHandler())
	for owner, provider := range appConfig.CsvProviders {
		if handler, err := app.LocalHandler(configDir, provider); err != nil {
			log.Panic(err, "register HTTP handler error")
		} else {
			r.GET("/local/"+owner, metrics, handler)
		}
	}
	r.GET("/metrics", app.MetricHandler())

	go func() {
		if err := r.Run(appConfig.BindAddress); err != nil {
			log.Error("start HTTP server error: %v", err.Error())
		}
	}()

	log.Info("start HTTP server success, bind address: %v", appConfig.BindAddress)

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}

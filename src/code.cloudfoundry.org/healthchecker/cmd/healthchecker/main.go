package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"time"

	"code.cloudfoundry.org/healthchecker/config"
	"code.cloudfoundry.org/healthchecker/watchdog"
	"code.cloudfoundry.org/lager/v3"
	"code.cloudfoundry.org/lager/v3/lagerflags"
)

const (
	SIGNAL_BUFFER_SIZE = 1024
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "", "Configuration File")
	flag.Parse()

	c, err := config.LoadConfig(configFile)
	if err != nil {
		panic(err)
	}

	logConfig := lagerflags.DefaultLagerConfig()
	logConfig.LogLevel = c.LogLevel
	logConfig.TimeFormat = lagerflags.FormatRFC3339
	logger, _ := lagerflags.NewFromConfig(c.ComponentName, logConfig)

	startupDelay := c.StartResponseDelayInterval + c.StartupDelayBuffer
	logger.Debug("Sleeping before gorouter responds to /health endpoint on startup", lager.Data{"sleep_time_seconds": startupDelay.Seconds()})
	time.Sleep(startupDelay)

	logger.Info("Starting")

	u := &url.URL{
		Scheme: c.HealthCheckEndpoint.Scheme,
		Host:   fmt.Sprintf("%s:%d", c.HealthCheckEndpoint.Host, c.HealthCheckEndpoint.Port),
		User:   url.UserPassword(c.HealthCheckEndpoint.User, c.HealthCheckEndpoint.Password),
		Path:   c.HealthCheckEndpoint.Path,
	}
	if c.HealthCheckEndpoint.Socket != "" {
		u.Opaque = c.HealthCheckEndpoint.Path
		u.Host = fmt.Sprintf("unix%s", c.HealthCheckEndpoint.Socket)
	}

	w := watchdog.NewWatchdog(u, c.ComponentName, c.FailureCounterFile, c.HealthCheckPollInterval, c.HealthCheckTimeout, logger)
	signals := make(chan os.Signal, SIGNAL_BUFFER_SIZE)
	signal.Notify(signals, watchdog.HandledSignals...)

	err = w.WatchHealthcheckEndpoint(context.Background(), signals)
	if err != nil {
		logger.Fatal("Error running healthcheck:", err)
	}
}

package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	boshdir "github.com/cloudfoundry/bosh-cli/director"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	boshuuid "github.com/cloudfoundry/bosh-utils/uuid"

	"github.com/cppforlife/bosh-stats/director"
	"github.com/cppforlife/bosh-stats/stats"
	"github.com/cppforlife/bosh-stats/stats/reporter"
)

const mainLogTag = "main"

var (
	debugOpt      = flag.Bool("debug", false, "Output debug logs")
	configPathOpt = flag.String("configPath", "", "Path to configuration file")
)

func main() {
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	logger, fs, _ := basicDeps(*debugOpt)
	defer logger.HandlePanic("Main")

	config, err := NewConfigFromPath(*configPathOpt, fs)
	ensureNoErr(logger, "Loading config", err)

	var dir boshdir.Director

	{
		directorFactory := director.NewFactory(config.Director, logger)

		dir, err = directorFactory.New()
		ensureNoErr(logger, "Failed building director", err)
	}

	eventsFactory := stats.NewEventsFactory(config.EventsStore.Path, fs)
	src := stats.NewDirector(dir, eventsFactory, logger)

	loggerReporter := reporter.NewLogger(logger)
	worker := stats.NewWorker(src, loggerReporter, logger)
	err = worker.Send()
	ensureNoErr(logger, "Sending stats", err)

	if config.Datadog.AppKey != "" {
		datadogReporter := reporter.NewDatadog(config.Datadog, logger)
		worker := stats.NewWorker(src, datadogReporter, logger)
		err = worker.Send()
		ensureNoErr(logger, "Sending stats to Datadog", err)
	}
}

func basicDeps(debug bool) (boshlog.Logger, boshsys.FileSystem, boshuuid.Generator) {
	logLevel := boshlog.LevelInfo

	// Debug generates a lot of log activity
	if debug {
		logLevel = boshlog.LevelDebug
	}

	logger := boshlog.NewWriterLogger(logLevel, os.Stderr, os.Stderr)
	fs := boshsys.NewOsFileSystem(logger)
	uuidGen := boshuuid.NewGenerator()
	return logger, fs, uuidGen
}

func ensureNoErr(logger boshlog.Logger, errPrefix string, err error) {
	if err != nil {
		logger.Error(mainLogTag, "%s: %s", errPrefix, err)
		os.Exit(1)
	}
}

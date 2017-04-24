package main

import (
	"encoding/json"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	"github.com/cppforlife/bosh-stats/director"
	"github.com/cppforlife/bosh-stats/stats/reporter"
)

type Config struct {
	Director    director.Config
	Datadog     reporter.DatadogConfig
	EventsStore EventsStore
}

type EventsStore struct {
	Path string
}

func NewConfigFromPath(path string, fs boshsys.FileSystem) (Config, error) {
	var config Config

	bytes, err := fs.ReadFile(path)
	if err != nil {
		return config, bosherr.WrapErrorf(err, "Reading config %s", path)
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, bosherr.WrapError(err, "Unmarshalling config")
	}

	err = config.Validate()
	if err != nil {
		return config, bosherr.WrapError(err, "Validating config")
	}

	return config, nil
}

func (c Config) Validate() error {
	err := c.Director.Validate()
	if err != nil {
		return bosherr.WrapError(err, "Validating 'Director' config")
	}

	err = c.Datadog.Validate()
	if err != nil {
		return bosherr.WrapError(err, "Validating 'Datadog' config")
	}

	if len(c.EventsStore.Path) == 0 {
		return bosherr.Error("Missing 'EventsStore.Path'")
	}

	return nil
}

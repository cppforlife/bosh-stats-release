package stats

import (
	"strconv"

	"github.com/cloudfoundry/bosh-cli/director"

	"github.com/cppforlife/bosh-stats/director/runtimeconfig"
)

type RuntimeConfig struct {
	director director.Director
}

func NewRuntimeConfig(director director.Director) RuntimeConfig {
	return RuntimeConfig{director}
}

func (f RuntimeConfig) Stats() ([]Stat, error) {
	rc, err := f.director.LatestRuntimeConfig()
	if err != nil {
		// todo check if runtime config is not found
		return nil, err
	}

	parsedRC, err := runtimeconfig.NewFromBytes([]byte(rc.Properties))
	if err != nil {
		return nil, err
	}

	var stats []Stat

	stats = append(stats, Stat{
		name:  "addons.count",
		value: strconv.Itoa(len(parsedRC.Addons)),
	})

	return stats, nil
}

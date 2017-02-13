package stats

import (
	boshdir "github.com/cloudfoundry/bosh-cli/director"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type Director struct {
	director      boshdir.Director
	eventsFactory EventsFactory

	logTag string
	logger boshlog.Logger
}

func NewDirector(director boshdir.Director, eventsFactory EventsFactory, logger boshlog.Logger) Director {
	return Director{director, eventsFactory, "stats.Director", logger}
}

func (f Director) Stats() ([]Stat, error) {
	var sources []Source

	sources = append(sources, NewDirectorInfo(f.director))
	sources = append(sources, NewCloudConfig(f.director))
	sources = append(sources, NewRuntimeConfig(f.director))
	sources = append(sources, NewReleases(f.director))
	sources = append(sources, NewStemcells(f.director))
	sources = append(sources, NewDeployments(f.director))
	sources = append(sources, f.eventsFactory.New(f.director))

	var allStats []Stat

	for _, src := range sources {
		stats, err := src.Stats()
		if err != nil {
			f.logger.Error(f.logTag, "Failed to get stats from '%T': %s", src, err)
		}

		if len(stats) > 0 {
			allStats = append(allStats, stats...)
		}
	}

	// todo augment with uuid, cpi, director version
	return allStats, nil
}

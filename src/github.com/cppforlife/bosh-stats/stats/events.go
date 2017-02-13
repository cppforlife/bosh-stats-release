package stats

import (
	boshdir "github.com/cloudfoundry/bosh-cli/director"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

type Events struct {
	storePath string

	fs       boshsys.FileSystem
	director boshdir.Director
}

func NewEvents(storePath string, fs boshsys.FileSystem, director boshdir.Director) Events {
	return Events{storePath, fs, director}
}

func (f Events) Stats() ([]Stat, error) {
	var stats []Stat

	_, err := f.director.Events(boshdir.EventsFilter{})
	if err != nil {
		return nil, err
	}

	return stats, nil
}

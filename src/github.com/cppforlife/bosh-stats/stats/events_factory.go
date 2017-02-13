package stats

import (
	boshdir "github.com/cloudfoundry/bosh-cli/director"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

type EventsFactory struct {
	storePath string
	fs        boshsys.FileSystem
}

func NewEventsFactory(storePath string, fs boshsys.FileSystem) EventsFactory {
	return EventsFactory{storePath, fs}
}

func (f EventsFactory) New(director boshdir.Director) Events {
	return NewEvents(f.storePath, f.fs, director)
}

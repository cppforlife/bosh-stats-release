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

	var prevCheckedID string
	var newCheckedID string

	if f.fs.FileExists(f.storePath) {
		content, err := f.fs.ReadFileString(f.storePath)
		if err != nil {
			return nil, err
		}

		prevCheckedID = content
	}

	var currPageTopID string

	for {
		events, err := f.director.Events(boshdir.EventsFilter{BeforeID: currPageTopID})
		if err != nil {
			return nil, err
		}

		if len(events) == 0 {
			break
		}
		currPageTopID = events[len(events)-1].ID()

		if len(newCheckedID) == 0 {
			newCheckedID = events[0].ID()
		}

		for _, ev := range events {
			if ev.ID() == prevCheckedID {
				break
			}

			if f.isDeploymentEnd(ev) && len(ev.Error()) > 0 {
				stats = append(stats, Stat{
					name:  "deployment.error",
					value: "1",
					tags: map[string]string{
						"error": ev.Error(),
					},
				})
			}
		}
	}

	err := f.fs.WriteFileString(f.storePath, newCheckedID)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (f Events) isDeploymentEnd(ev boshdir.Event) bool {
	return ev.Action() == "create" && ev.ObjectType() == "deployment" && len(ev.ParentID()) > 0
}

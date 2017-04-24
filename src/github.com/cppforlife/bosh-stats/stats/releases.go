package stats

import (
	"strconv"

	boshdir "github.com/cloudfoundry/bosh-cli/director"
)

type Releases struct {
	director boshdir.Director
}

func NewReleases(director boshdir.Director) Releases {
	return Releases{director}
}

func (f Releases) Stats() ([]Stat, error) {
	releases, err := f.director.Releases()
	if err != nil {
		return nil, err
	}

	var stats []Stat
	directorTags := f.directorTags()

	stats = append(stats, Stat{
		name:  "releases.count",
		value: strconv.Itoa(len(releases)),
		tags:  directorTags,
	})

	for _, rel := range releases {
		tags := map[string]string{
			"name":    rel.Name(),
			"version": rel.Version().String(),
		}

		stat := Stat{
			name:  "release",
			value: "1",
			tags:  merge(directorTags, tags),
			// todo used?
		}

		stats = append(stats, stat)
	}

	return stats, nil
}

func (f Releases) directorTags() map[string]string {
	info, err := f.director.Info()
	if err != nil {
		// TODO: handle/return error
		return nil
	}

	tags := make(map[string]string)

	tags["director.name"] = info.Name
	tags["director.uuid"] = info.UUID

	return tags
}

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

	stats = append(stats, Stat{
		name:  "releases.count",
		value: strconv.Itoa(len(releases)),
	})

	for _, rel := range releases {
		stat := Stat{
			name:  "release",
			value: "true",
			tags: map[string]string{
				"name":    rel.Name(),
				"version": rel.Version().String(),
				// todo used?
			},
		}

		stats = append(stats, stat)
	}

	return stats, nil
}

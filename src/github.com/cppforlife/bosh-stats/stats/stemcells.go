package stats

import (
	"strconv"

	boshdir "github.com/cloudfoundry/bosh-cli/director"
)

type Stemcells struct {
	director boshdir.Director
}

func NewStemcells(director boshdir.Director) Stemcells {
	return Stemcells{director}
}

func (f Stemcells) Stats() ([]Stat, error) {
	stemcells, err := f.director.Stemcells()
	if err != nil {
		return nil, err
	}

	var stats []Stat
	directorTags := f.directorTags()

	// name="stemcell" value="true" tags="map[string]string{"version":"3363.19", "name":"bosh-warden-boshlite-ubuntu-trusty-go_agent"}"
	// name="stemcells.count" value="1" tags="map[string]string(nil)"

	stats = append(stats, Stat{
		name:  "stemcells.count",
		value: strconv.Itoa(len(stemcells)),
		tags:  directorTags,
	})

	for _, stem := range stemcells {
		tags := map[string]string{
			"name":    stem.Name(),
			"version": stem.Version().String(),
		}

		stat := Stat{
			name:  "stemcell",
			value: "1",
			tags:  merge(directorTags, tags),
			// todo used?
		}

		stats = append(stats, stat)
	}

	return stats, nil
}

func (f Stemcells) directorTags() map[string]string {
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

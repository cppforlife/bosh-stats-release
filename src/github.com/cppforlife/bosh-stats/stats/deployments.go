package stats

import (
	"strconv"

	boshdir "github.com/cloudfoundry/bosh-cli/director"
)

type Deployments struct {
	director boshdir.Director
}

func NewDeployments(director boshdir.Director) Deployments {
	return Deployments{director}
}

func (f Deployments) Stats() ([]Stat, error) {
	deps, err := f.director.Deployments()
	if err != nil {
		return nil, err
	}

	stats := []Stat{}
	directorTags := f.directorTags()

	stats = append(stats, Stat{
		name:  "deployments.count",
		value: strconv.Itoa(len(deps)),
		tags:  directorTags,
	})

	var allInstances []boshdir.Instance

	for _, dep := range deps {
		insts, err := dep.Instances()
		if err != nil {
			return stats, nil
		}

		tags := directorTags
		tags["deployment.name"] = dep.Name()

		stats = append(stats, Stat{
			name:  "deployment.instances.count",
			value: strconv.Itoa(len(insts)),
			tags:  tags,
		})

		allInstances = append(allInstances, insts...)
	}

	stats = append(stats, Stat{
		name:  "instances.count",
		value: strconv.Itoa(len(allInstances)),
		tags:  directorTags,
	})

	return stats, nil
}

func (f Deployments) directorTags() map[string]string {
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

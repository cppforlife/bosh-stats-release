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

	var stats []Stat

	stats = append(stats, Stat{
		name:  "deployments.count",
		value: strconv.Itoa(len(deps)),
	})

	var allInstances []boshdir.Instance

	for _, dep := range deps {
		insts, err := dep.Instances()
		if err != nil {
			return stats, nil
		}

		stats = append(stats, Stat{
			name:  "deployment.instances.count",
			value: strconv.Itoa(len(insts)),
		})

		allInstances = append(allInstances, insts...)
	}

	stats = append(stats, Stat{
		name:  "instances.count",
		value: strconv.Itoa(len(allInstances)),
	})

	return stats, nil
}

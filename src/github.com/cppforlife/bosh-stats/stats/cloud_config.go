package stats

import (
	"strconv"
	"strings"

	"github.com/cloudfoundry/bosh-cli/director"

	"github.com/cppforlife/bosh-stats/director/cloudconfig"
)

type CloudConfig struct {
	director director.Director
}

func NewCloudConfig(director director.Director) CloudConfig {
	return CloudConfig{director}
}

func (f CloudConfig) Stats() ([]Stat, error) {
	cc, err := f.director.LatestCloudConfig()
	if err != nil {
		// todo check if cloud config is not found
		return nil, err
	}

	parsedCC, err := cloudconfig.NewFromBytes([]byte(cc.Properties))
	if err != nil {
		return nil, err
	}

	stats := []Stat{f.azNames(parsedCC)}

	stats = append(stats, Stat{
		name:  "azs.count",
		value: strconv.Itoa(len(parsedCC.AZs)),
	})

	stats = append(stats, Stat{
		name:  "disk_types.count",
		value: strconv.Itoa(len(parsedCC.DiskTypes)),
	})

	stats = append(stats, Stat{
		name:  "vm_types.count",
		value: strconv.Itoa(len(parsedCC.VMTypes)),
	})

	stats = append(stats, Stat{
		name:  "networks.count",
		value: strconv.Itoa(len(parsedCC.Networks)),
	})

	stats = append(stats, Stat{
		name:  "networks.manual.count",
		value: strconv.Itoa(f.networksWithType(parsedCC, "manual") + f.networksWithType(parsedCC, "")),
	})

	stats = append(stats, Stat{
		name:  "networks.dynamic.count",
		value: strconv.Itoa(f.networksWithType(parsedCC, "dynamic")),
	})

	stats = append(stats, Stat{
		name:  "networks.vip.count",
		value: strconv.Itoa(f.networksWithType(parsedCC, "vip")),
	})

	stats = append(stats, Stat{
		name:  "compilation.workers",
		value: strconv.Itoa(parsedCC.Compilation.Workers),
	})

	return stats, nil
}

func (f CloudConfig) azNames(parsedCC cloudconfig.Schema) Stat {
	var azNames []string

	for _, az := range parsedCC.AZs {
		azNames = append(azNames, az.Name)
	}

	return Stat{
		name:  "azs.names",
		value: strings.Join(azNames, ","),
	}
}

func (f CloudConfig) networksWithType(parsedCC cloudconfig.Schema, netType string) int {
	var count int
	for _, net := range parsedCC.Networks {
		if net.Type == netType {
			count++
		}
	}
	return count
}

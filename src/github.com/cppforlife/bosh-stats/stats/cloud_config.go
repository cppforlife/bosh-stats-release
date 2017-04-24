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

	stats := []Stat{}
	tags := f.directorTags()

	stats = append(stats, Stat{
		name:  "azs",
		value: strconv.Itoa(len(parsedCC.AZs)),
		tags:  merge(tags, f.azTags(parsedCC)),
	})

	stats = append(stats, Stat{
		name:  "azs.count",
		value: strconv.Itoa(len(parsedCC.AZs)),
		tags:  tags,
	})

	stats = append(stats, Stat{
		name:  "disk_types.count",
		value: strconv.Itoa(len(parsedCC.DiskTypes)),
		tags:  tags,
	})

	stats = append(stats, Stat{
		name:  "vm_types.count",
		value: strconv.Itoa(len(parsedCC.VMTypes)),
		tags:  tags,
	})

	stats = append(stats, Stat{
		name:  "networks.count",
		value: strconv.Itoa(len(parsedCC.Networks)),
		tags:  tags,
	})

	stats = append(stats, Stat{
		name:  "networks.manual.count",
		value: strconv.Itoa(f.networksWithType(parsedCC, "manual") + f.networksWithType(parsedCC, "")),
		tags:  tags,
	})

	stats = append(stats, Stat{
		name:  "networks.dynamic.count",
		value: strconv.Itoa(f.networksWithType(parsedCC, "dynamic")),
		tags:  tags,
	})

	stats = append(stats, Stat{
		name:  "networks.vip.count",
		value: strconv.Itoa(f.networksWithType(parsedCC, "vip")),
		tags:  tags,
	})

	stats = append(stats, Stat{
		name:  "compilation.workers",
		value: strconv.Itoa(parsedCC.Compilation.Workers),
		tags:  tags,
	})

	return stats, nil
}

func (f CloudConfig) azTags(parsedCC cloudconfig.Schema) map[string]string {
	tags := make(map[string]string)
	var azTags []string

	for _, az := range parsedCC.AZs {
		azTags = append(azTags, az.Name)
	}

	tags["names"] = strings.Join(azTags, ",")
	return tags
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

func (f CloudConfig) directorTags() map[string]string {
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

func merge(one map[string]string, two map[string]string) map[string]string {
	result := make(map[string]string)

	for k, v := range one {
		result[k] = v
	}

	for k, v := range two {
		result[k] = v
	}

	return result
}

package stats

import (
	boshdir "github.com/cloudfoundry/bosh-cli/director"
)

type DirectorInfo struct {
	director boshdir.Director
}

func NewDirectorInfo(director boshdir.Director) DirectorInfo {
	return DirectorInfo{director}
}

func (f DirectorInfo) Stats() ([]Stat, error) {
	info, err := f.director.Info()
	if err != nil {
		return nil, err
	}

	var stats []Stat

	tags := make(map[string]string)
	tags["name"] = info.Name // todo sensitive info
	tags["uuid"] = info.UUID
	tags["version"] = info.Version
	tags["auth.type"] = info.Auth.Type
	tags["auth.cpi"] = info.CPI

	stats = append(stats, Stat{
		name:  "director.info",
		value: "1",
		tags:  tags,
	})

	// todo features

	return stats, nil
}

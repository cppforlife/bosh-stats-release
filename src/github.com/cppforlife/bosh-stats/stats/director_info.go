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

	stats = append(stats, Stat{
		name:  "director.version",
		value: info.Version,
	})

	stats = append(stats, Stat{
		name:  "director.uuid",
		value: info.UUID,
	})

	stats = append(stats, Stat{
		name:  "director.auth.type",
		value: info.Auth.Type,
	})

	stats = append(stats, Stat{
		name:  "director.cpi",
		value: info.CPI,
	})

	// todo features

	return stats, nil
}

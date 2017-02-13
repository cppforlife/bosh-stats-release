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

	stats = append(stats, Stat{
		name:  "stemcells.count",
		value: strconv.Itoa(len(stemcells)),
	})

	for _, stem := range stemcells {
		stat := Stat{
			name:  "stemcell",
			value: "true",
			tags: map[string]string{
				"name":    stem.Name(),
				"version": stem.Version().String(),
				// todo used?
			},
		}

		stats = append(stats, stat)
	}

	return stats, nil
}

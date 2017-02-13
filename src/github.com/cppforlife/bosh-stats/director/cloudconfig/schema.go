package cloudconfig

import (
	"gopkg.in/yaml.v2"
)

type Schema struct {
	AZs      []AZ
	Networks []Network

	DiskTypes []DiskType `yaml:"disk_types"`
	VMTypes   []VMType   `yaml:"vm_types"`

	Compilation Compilation
}

type AZ struct {
	Name            string
	CloudProperties interface{} `yaml:"cloud_properties"`
}

type Network struct {
	Name string
	Type string
}

type DiskType struct {
	Name string
}

type VMType struct {
	Name string
}

type Compilation struct {
	Workers int
}

func NewFromBytes(bytes []byte) (Schema, error) {
	var schema Schema

	err := yaml.Unmarshal(bytes, &schema)
	if err != nil {
		return Schema{}, err
	}

	return schema, nil
}

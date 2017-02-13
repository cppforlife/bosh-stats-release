package runtimeconfig

import (
	"gopkg.in/yaml.v2"
)

type Schema struct {
	Addons []Addon
}

type Addon struct {
	Name string
}

func NewFromBytes(bytes []byte) (Schema, error) {
	var schema Schema

	err := yaml.Unmarshal(bytes, &schema)
	if err != nil {
		return Schema{}, err
	}

	return schema, nil
}

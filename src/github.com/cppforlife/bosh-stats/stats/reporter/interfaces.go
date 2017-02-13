package reporter

type Stat interface {
	Name() string
	Value() string
	Tags() map[string]string
}

type Reporter interface {
	Report([]Stat) error
}

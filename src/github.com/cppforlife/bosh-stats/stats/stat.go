package stats

type Stat struct {
	name  string
	value string
	tags  map[string]string
}

func (s Stat) Name() string            { return s.name }
func (s Stat) Value() string           { return s.value }
func (s Stat) Tags() map[string]string { return s.tags }

package stats

type Source interface {
	Stats() ([]Stat, error)
}

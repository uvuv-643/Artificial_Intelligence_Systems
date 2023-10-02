package data

type LinearNorm struct {
	mins   map[string]float64
	maxs   map[string]float64
	target string
}

func (ln *LinearNorm) Fit(dataset Dataset, target string) {
	ln.mins = Min(dataset)
	ln.maxs = Max(dataset)
	ln.target = target
}

func (ln *LinearNorm) Apply(dataset Dataset) Dataset {
	newDataset := copyDataset(dataset)
	for _, field := range dataset.Fields() {
		if field == ln.target {
			continue
		}
		for _, row := range newDataset.Data {
			row[field] = (row[field] - ln.mins[field]) / (ln.maxs[field] - ln.mins[field])
		}
	}
	return newDataset
}

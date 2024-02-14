package cmd

type Data struct {
	Key   string
	Value int
}

type Repository struct {
	DataHM map[string]Data
}

func NewRepository() *Repository {
	return &Repository{make(map[string]Data)}
}

func (r *Repository) ConsultaChave(key string) (bool, int, error) {
	item := r.DataHM[key]
	if item.Key == "" {
		return false, 0, nil
	}
	return true, item.Value, nil
}

func (r *Repository) IncluiChave(key string, value int) (Data, error) {
	item := Data{key, value}
	r.DataHM[item.Key] = item
	return item, nil
}

func (r *Repository) Incrementa(key string) (int, error) {
	item := r.DataHM[key]
	if item.Key == "" {
		item, _ = r.IncluiChave(key, 1)
	} else {
		item.Value = item.Value + 1
		r.DataHM[item.Key] = item
	}
	return item.Value, nil
}

func (r *Repository) ExcluiChave(key string) error {
	delete(r.DataHM, key)
	return nil
}

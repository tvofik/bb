package main

type Translation struct {
	name         string
	abbreviation string
}

func (t Translation) FilterValue() string {
	return t.name
}
func (t Translation) Title() string {
	return t.abbreviation
}

func (t Translation) Description() string {
	return t.name
}

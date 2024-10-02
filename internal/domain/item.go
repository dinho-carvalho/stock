package domain

type Item struct {
	id   int
	name string
}

func NewItem(name string) Item {
	return Item{
		name: name,
	}
}

func (i *Item) Name() string {
	return i.name
}

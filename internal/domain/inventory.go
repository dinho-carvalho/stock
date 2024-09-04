package domain

type Inventory struct {
	id int
	name string
}

func NewInventory(name string) Inventory {
	return Inventory {
		name: name
	}
}

func (i *Inventory) Name() string {
	return i.name
}

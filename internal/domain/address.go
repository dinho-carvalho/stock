package domain

type Address struct {
	name string
}

func NewAddress(name string) *Address {
	return &Address{name: name}
}

func (a *Address) Name() string {
	return a.name
}

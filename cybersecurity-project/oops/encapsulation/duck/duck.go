package duck

type duck struct {
	name string //encapsulation of private variable
}

func New(name string) *duck {
	return &duck{
		name: name,
	}
}
func (d *duck) GetName() string {
	return d.name
}

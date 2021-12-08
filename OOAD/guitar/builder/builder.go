package builder

type Builder int

const (
	Fender = 1
	Martin = 2
	Gibson = 3
)

func (b Builder) String() string {
	if b == 0 {
		return ""
	}
	return [...]string{"Fender", "Martin", "Gibson"}[b-1]
}

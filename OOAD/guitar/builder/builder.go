package builder

type Builder int

const (
	Fender Builder = 1
	Martin Builder = 2
	Gibson Builder = 3
)

func (b Builder) String() string {
	if b == 0 {
		return ""
	}
	return [...]string{"Fender", "Martin", "Gibson"}[b-1]
}

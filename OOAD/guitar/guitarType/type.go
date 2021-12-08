package guitarType

type GuitarType int

const (
	Accoustic GuitarType = 101
	Electric  GuitarType = 201
)

func (t GuitarType) String() string {
	if t == 0 {
		return ""
	}
	s := make([]string, 203)
	s[100] = "Accoustic"
	s[200] = "Electric"
	return s[t-1]
}

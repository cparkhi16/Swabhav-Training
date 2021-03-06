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
	v := []string{"", "Accoustic", "Electric"}
	if t == 101 || t == 201 {
		index := (t - 1) / 100
		return v[index]
	}
	return v[0]
}

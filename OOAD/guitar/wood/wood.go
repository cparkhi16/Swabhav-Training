package wood

type Wood int

const (
	Mahagony = 11
	Maple    = 21
	Cocobolo = 31
	Cedar    = 41
	Sitka    = 51
)

func (w Wood) String() string {
	if w == 0 {
		return ""
	}
	s := make([]string, 52)
	s[10] = "Mahagony"
	s[20] = "Maple"
	s[30] = "Cocobolo"
	s[40] = "Cedar"
	s[50] = "Sitka"
	return s[w-1]
}

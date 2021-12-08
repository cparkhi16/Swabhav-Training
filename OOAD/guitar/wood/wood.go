package wood

type Wood int

const (
	Mahagony Wood = 11
	Maple    Wood = 21
	Cocobolo Wood = 31
	Cedar    Wood = 41
	Sitka    Wood = 51
)

func (w Wood) String() string {
	/*if w == 0 {
		return ""
	}
	s := make([]string, 52)
	s[10] = "Mahagony"
	s[20] = "Maple"
	s[30] = "Cocobolo"
	s[40] = "Cedar"
	s[50] = "Sitka"*/
	v := []string{"", "Mahagony", "Maple", "Cocobolo", "Cedar", "Sitka"}
	if w == 11 || w == 21 || w == 31 || w == 41 || w == 51 {
		index := (w - 1) / 10
		return v[index]
	}
	return v[0]

}

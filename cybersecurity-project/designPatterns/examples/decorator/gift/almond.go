package gift

type almond struct {
	from string
}

func NewAlmond(from string) almond {
	return almond{
		from: from,
	}
}

func (a almond) WhatIsThis() string {
	return "almonds from " + a.from
}

package gift

type cashew struct {
	from string
}

func NewCashew(from string) cashew {
	return cashew{
		from: from,
	}
}

func (c cashew) WhatIsThis() string {
	return "cashew from " + c.from
}

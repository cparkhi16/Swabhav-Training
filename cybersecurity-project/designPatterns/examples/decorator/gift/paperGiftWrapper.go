package gift

type glitterPaperWrapper struct {
	giftWrapped gift
}

func NewGlitterPaperWrapper(giftWrapped gift, from string) glitterPaperWrapper {
	return glitterPaperWrapper{
		giftWrapped: giftWrapped,
	}
}

func (g glitterPaperWrapper) WhatIsThis() string {
	return "plane paper wrapped gift box with " + g.giftWrapped.WhatIsThis()
}

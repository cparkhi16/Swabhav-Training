package gift

type glitterGiftWrapper struct {
	giftWrapped gift
}

func NewGlitterGiftWrapper(giftWrapped gift, from string) glitterGiftWrapper {
	return glitterGiftWrapper{
		giftWrapped: giftWrapped,
	}
}

func (g glitterGiftWrapper) WhatIsThis() string {
	return "glittery shiny gift box with " + g.giftWrapped.WhatIsThis()
}

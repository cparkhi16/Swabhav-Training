package user

import "day_8/hotel"

type User struct {
	userName    string
	userContact string
	orderList   []hotel.Order
}

func (u *User) NewUser(userName, contact string) *User {
	return &User{
		userName:    userName,
		userContact: contact,
	}
}
func (u *User) AddOrderForUser(o hotel.Order) *User {
	u.orderList = append(u.orderList, o)
	return u
}

package hotel

import "fmt"

type User struct {
	userName    string
	userContact string
	orderList   []Order
}

func NewUser(userName, contact string) *User {
	return &User{
		userName:    userName,
		userContact: contact,
	}
}
func (u *User) AddOrderForUser(o Order, ch <-chan string) {
	OrderStatus := <-ch
	fmt.Printf("For user :%v,Order ID : %v %v\n", u.userName, o.orderID, OrderStatus)
	u.orderList = append(u.orderList, o)
}
func (u *User) GetOrder() {
	fmt.Println("Your order is ready--------------- ", u.userName)
	fmt.Println("User name ", u.userName)
	fmt.Println("User contact", u.userContact)
	fmt.Println("User order list ", u.orderList)
}
func (u *User) AssignDeskForUser(ch <-chan int) {
	serve := make(chan struct{})
	go u.DistributeOrder(serve)
	serve <- struct{}{}
	DeskNumber := <-ch
	fmt.Printf("%v has been assigned %v desk Number \n", u.userName, DeskNumber)
}
func (u *User) InvoiceForUser(ch <-chan struct{}) {
	fmt.Printf("Invoice generation started for user -- %v\n", u.userName)
	var totalAmount uint
	for _, val := range u.orderList {
		for _, v := range val.items {
			totalAmount = totalAmount + v.itemPrice
		}
	}
	fmt.Printf("Total amount for user %v is %v\n", u.userName, totalAmount)
	<-ch
}
func (u *User) DistributeOrder(c <-chan struct{}) {
	fmt.Printf("Distributing Order to %v \n", u.userName)
	<-c
}

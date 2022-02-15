package main

import "fmt"

type account struct{
	accountNo int
}

func newAccount(accountNo int)*account{
	return &account{
		accountNo:accountNo,
	}
}

func(a *account) checkAccount(accountNo int){
	if a.accountNo==accountNo{
		fmt.Println("account number is equal")
	}else{
		fmt.Println("not equal")
	}
}

type securityCode struct{
	code int
}

func newCode(code int)*securityCode{
	return &securityCode{
		code:code,
	}
}

func(s *securityCode) checkCode(securityCode int){
	if s.code==securityCode{
		fmt.Println("code is equal")
	}else{
		fmt.Println("not equal")
	}
}

type wallet struct{
	balance int
}

func newWallet(balance int)*wallet{
	return &wallet{
		balance:balance,
	}
}

func(w *wallet)credit(amount int){
	w.balance=w.balance+amount
}

func(w *wallet)debit(amount int){
	w.balance=w.balance-amount
}

type ledger struct{
	entries []string
}

func(l *ledger) makeEntry(operation string){
	l.entries=append(l.entries,operation)
}

type notification struct{
	msg string
}

func(n *notification) sendNotification(msg string){
	n.msg=msg
	fmt.Println("Notification:-",msg)
}

func main(){
	shanAcc:=newAccount(123)
	shanSecurityCode:=newCode(333)
	shanWallet:=newWallet(2000)
	var str=make([]string,10)
	bankLedger:=ledger{
		entries:str,
	}
	notification:=notification{
		msg:"",
	}
	shanAcc.checkAccount(123)
	shanSecurityCode.checkCode(333)
	shanWallet.credit(1000)
	bankLedger.makeEntry("money transfer of 1000 to shanAcc")
	notification.sendNotification("money transfer of 1000 to shanAcc")
}
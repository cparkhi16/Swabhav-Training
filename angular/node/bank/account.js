const EventEmitter = require('events');

class Account extends EventEmitter{
    fields = {
        accountNumber:0,
        balance:0,
    };
    
    constructor(accountNumber,balance){
        super();
        this.accountNumber=accountNumber;
        this.balance=balance;
    }
    creditMoney(money) {
        this.balance=this.balance+money;
        //s=money+"Credited"
       // this.emit('event',"Money credited ->",money)
       this.updateLedger("Money credited ->",money,"  Account Number ",this.accountNumber)
    }
      
    debitMoney(money){
        console.log(this.balance)
        this.balance=this.balance-money
       // this.emit('event',"Money debited ->",money)
       this.updateLedger("Money dedited ->",money,"  Account Number ",this.accountNumber)
       console.log(this.balance)
    }
    updateLedger(update,money,acc,accNumber){
        this.emit('event',update,money,acc,accNumber)
    }
}
module.exports=Account
const EventEmitter=require('events')

class AccountEventEmitter extends EventEmitter{
    balanceChange(accountNo,amount,transactionType){
        this.emit("balanceChange",{accountNo,amount,transactionType});
    }
}

accountEventEmitter=new AccountEventEmitter();

class Account{
    constructor(accountNo,balance){
        this.accountNo=accountNo;
        this.balance=balance;
    }

    deposit(amount){
        this.balance=this.balance+amount;
        accountEventEmitter.balanceChange(this.accountNo,amount,"deposit");
    }

    withdraw(amount){
        this.balance=this.balance-amount;
        accountEventEmitter.balanceChange(this.accountNo,amount,"withdraw");
    }

    getBalance(){
        return this.balance;
    }

}

module.exports.Account=Account
module.exports.AccountEventEmitter=AccountEventEmitter
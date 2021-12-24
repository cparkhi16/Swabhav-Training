const Account=require('./account')
const Ledger=require('./ledger')
class User {

    constructor(name,email,account) {
      this.name = name;
      this.email=email;
      this.accounts=account;
    }
    getUserBalance(accountNumber){
        for (let i = 0; i <this.accounts.length; i++) {
            if (this.accounts[i].accountNumber==accountNumber){
                return this.accounts[i].balance
            }
          }
    }
  }
johnFirstAccount= new Account(123,5000)
johnSecondAccount=new Account(124,500)
johnLedger=new Ledger(johnFirstAccount)

let user = new User("John","j@fp.com",[johnFirstAccount,johnSecondAccount]);
johnLedger.BalanceChange()
johnFirstAccount.creditMoney(5000)
b=user.getUserBalance(123)
console.log("Balance of john's account ",b)

johnFirstAccount.debitMoney(1000)
v=user.getUserBalance(123)
console.log("Balance of john's account ",v)
//module.exports=johnAccount
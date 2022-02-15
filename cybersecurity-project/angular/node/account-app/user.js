const account=require('./account')

class User{
    
    constructor(name){
        this.name=name;
        this.accountList=new Map()
        //this.accountList.set(accountNo,new Account(accountNo,balance));
    }

    getAccount(accountNo){
        //console.log(this.accountList.get(accountNo));
        return this.accountList.get(accountNo);
    }

    addAccount(accountNo,balance){
        this.accountList.set(accountNo,new account.Account(accountNo,balance));
        //this.accountList[accountNo]=new Account(accountNo,balance)
    }
}

module.exports=User
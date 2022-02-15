const fs=require('fs')

class Ledger{
    constructor(accountEventEmitter){
        this.accountEventEmitter=accountEventEmitter;
    }

    listenToBalanceChange(){
        accountEventEmitter.on("balanceChange",function(a){
            console.log("Event caught",a);
            //writefile
            fs.writeFile('./file.txt',JSON.stringify(a)+'\n',{
                encoding: "utf8",
                flag: "a",
                mode: 0o666
                },(err)=>{
                if(err){
                    console.log(err);
                }
            });
        })        
    }
}

module.exports.Ledger=Ledger;
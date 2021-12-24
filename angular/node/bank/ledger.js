const fs = require("fs");

class Ledger{
   constructor(account){
       this.account=account;
   }

   BalanceChange(){
      this.account.on('event', (a,b,c,d) => {
         console.log(a,b);
         fs.appendFile('input.txt', a+b+c+d+" ", function(err) {
           if (err) {
              return console.error(err);
           }
           console.log("Data written successfully!");
           console.log("Let's read newly written data");
        
           fs.readFile('input.txt', function (err, data) {
              if (err) {
                 return console.error(err);
              }
              console.log("Asynchronous read: " + data.toString());
           });
        });
       });
       
   }
}

module.exports=Ledger;

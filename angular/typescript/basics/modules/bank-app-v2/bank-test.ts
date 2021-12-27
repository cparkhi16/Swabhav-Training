import xyz,{Account} from "./bank-module";
import * as BankModule from "./bank-module"

var bank=new xyz("ICICI","Dom")
console.log(bank)

var c =new Account(78,"23n2")
console.log(c)

var newBank= new BankModule.default("DNS","THA")
console.log(newBank)

var newAcc=new BankModule.Account(231,"hwd")
console.log(newAcc)

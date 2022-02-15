//import xyz,{Customer,Account} from "./bank-module"
//import {Customer,Account} from "./bank-module"
import * as bank from "./bank-module"

let b1=new bank.default("x","x")
console.log("bank-",b1)
let c1=new bank.Customer("r","r")
console.log("customer-",c1)
let a1=new bank.Account("e","e")
console.log("account-",a1)
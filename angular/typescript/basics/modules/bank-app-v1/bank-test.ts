import {Bank,Account,Customer} from "./bank-module";

var b =new Bank("Axis","Thane")
var a=new Account(1234,"1234de")
var c =new Customer("Chinmay",21)
console.log({b,a,c})
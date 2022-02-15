"use strict";
exports.__esModule = true;
var bank_module_1 = require("./bank-module");
var b1 = new bank_module_1.Bank("x", "x");
console.log("bank-", b1);
var c1 = new bank_module_1.Customer("r", "r");
console.log("customer-", c1);
var a1 = new bank_module_1.Account("e", "e");
console.log("account-", a1);

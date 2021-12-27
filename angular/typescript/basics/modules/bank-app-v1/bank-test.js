"use strict";
exports.__esModule = true;
var bank_module_1 = require("./bank-module");
var b = new bank_module_1.Bank("Axis", "Thane");
var a = new bank_module_1.Account(1234, "1234de");
var c = new bank_module_1.Customer("Chinmay", 21);
console.log({ b: b, a: a, c: c });

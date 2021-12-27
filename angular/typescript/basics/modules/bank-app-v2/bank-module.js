"use strict";
exports.__esModule = true;
exports.Customer = exports.Account = void 0;
var Account = /** @class */ (function () {
    function Account(accNumber, ifsc) {
        this.accNumber = accNumber;
        this.ifsc = ifsc;
    }
    return Account;
}());
exports.Account = Account;
var Bank = /** @class */ (function () {
    function Bank(Name, Location) {
        this.Name = Name;
        this.Location = Location;
    }
    return Bank;
}());
exports["default"] = Bank;
var Customer = /** @class */ (function () {
    function Customer(Name, age) {
        this.Name = Name;
        this.age = age;
    }
    return Customer;
}());
exports.Customer = Customer;

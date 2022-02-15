"use strict";
exports.__esModule = true;
exports.Customer = exports.Account = exports.Bank = void 0;
var Bank = /** @class */ (function () {
    function Bank(name, location) {
        this.name = name;
        this.location = location;
    }
    return Bank;
}());
exports.Bank = Bank;
var Account = /** @class */ (function () {
    function Account(no, ifsc) {
        this.no = no;
        this.ifsc = ifsc;
    }
    return Account;
}());
exports.Account = Account;
var Customer = /** @class */ (function () {
    function Customer(name, address) {
        this.name = name;
        this.address = address;
    }
    return Customer;
}());
exports.Customer = Customer;

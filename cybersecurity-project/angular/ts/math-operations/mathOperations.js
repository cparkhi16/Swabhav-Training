var add = /** @class */ (function () {
    function add(num1, num2) {
        this.num1 = num1;
        this.num2 = num2;
    }
    add.prototype.operation = function () {
        return this.num1 + this.num2;
    };
    return add;
}());
var multiply = /** @class */ (function () {
    function multiply(num1, num2) {
        this.num1 = num1;
        this.num2 = num2;
    }
    multiply.prototype.operation = function () {
        return this.num1 * this.num2;
    };
    return multiply;
}());
function performOperation(m) {
    console.log(m.operation());
}
function subtract(num1, num2) {
    return num1 - num2;
}
function divide(num1, num2) {
    return num1 / num2;
}
function performOperation2(num1, num2, func) {
    return func(num1, num2);
}
var a = new add(2, 3);
performOperation(a);
console.log(performOperation2(8, 5, subtract));

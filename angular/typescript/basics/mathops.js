function mathOperations(callback, num1, num2) {
    return callback(num1, num2);
}
function Add(a, b) {
    return a + b;
}
function Subtract(a, b) {
    return a - b;
}
function Divide(a, b) {
    if (b == 0) {
        return -1;
    }
    return a / b;
}
function Multiply(a, b) {
    return a * b;
}
var add = mathOperations(Add, 5, 4);
console.log("Addition ", add);
var sub = mathOperations(Subtract, 5, 4);
console.log("Subtraction ", sub);
var mul = mathOperations(Multiply, 5, 4);
console.log("Multiplication ", mul);
var div = mathOperations(Divide, 5, 4);
console.log("Division ", div);

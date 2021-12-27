class MathOps {
    mathOperation(num1, num2, operation) {
        return operation(num1, num2);
    }
}
function Add(num1, num2) {
    return num1 + num2;
}
function Sub(num1, num2) {
    return num1 - num2;
}
var m = new MathOps();
console.log(m.mathOperation(5, 5, Add));
console.log(m.mathOperation(5, 4, Sub));

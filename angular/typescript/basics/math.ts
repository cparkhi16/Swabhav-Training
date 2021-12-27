interface Calculator{
    mathOperation(num1:number,num2:number,operation:(num1:number,num2:number)=>number):number;
}

class MathOps implements Calculator{
    mathOperation(num1:number,num2:number,operation:(num1:number,num2:number)=>number){
        return operation(num1,num2);
    }
}

function Add(num1:number,num2:number){
    return num1+num2;
}
function Sub(num1:number,num2:number){
    return num1-num2;
}

var m =new MathOps()
console.log(m.mathOperation(5,5,Add))
console.log(m.mathOperation(5,4,Sub))
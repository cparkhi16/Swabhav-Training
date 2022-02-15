interface mathOperation{
    num1:number;
    num2:number;
    operation():number;
}

class add{
    public num1:number
    public num2:number
    constructor(num1,num2){
        this.num1=num1;
        this.num2=num2;
    }
    operation(){
        return this.num1+this.num2;
    }
}

class multiply{
    public num1:number
    public num2:number
    constructor(num1,num2){
        this.num1=num1;
        this.num2=num2;
    }
    operation(){
        return this.num1*this.num2;
    }
}

function performOperation(m:mathOperation){
    console.log(m.operation())
}

function subtract(num1:number,num2:number){
    return num1-num2;
}

function divide(num1:number,num2:number){
    return num1/num2;
}

function performOperation2(num1:number,num2:number,func:(x: number, y: number) => number){
    return func(num1,num2);
}
var a=new add(2,3)
performOperation(a)

console.log(performOperation2(8,5,subtract))

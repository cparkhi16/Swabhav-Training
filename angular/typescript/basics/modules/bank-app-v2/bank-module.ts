 export class Account{
    constructor(public accNumber:number,public ifsc:string){

    }
}

export default class Bank{
    constructor(public Name:string,public Location:string){
        
    }
}
export class Customer{
    constructor(public Name:string,public age:number){
        
    }
}
interface IPerson{
    firstName:string
    lastName:String
    age?:number
    getFirstName(): string; 
}

class User{
    firstName:string
    lastName:String
    age?:number
    constructor(firstName:string,lastName:string,age?:number){
        this.firstName=firstName;
        this.lastName=lastName;
        this.age=age;
    }

    getFirstName(){
        return this.firstName
    }
}

function showDetails(people:IPerson[]){
    for(var person of people){
        console.log(person.firstName,person.lastName,person.age)
    }
}

var shan=new User("shan","a",64)
console.log(shan.getFirstName())
var sumi=new User("sumi","b")
var people=[shan,sumi]
showDetails(people)



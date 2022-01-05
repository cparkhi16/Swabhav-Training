import { Component, OnInit } from '@angular/core';


class User{
  firstName:string
  age:number
  constructor(fN:string,a:number){
    this.firstName=fN
    this.age=a
  }
}
class Student{
  srno:number
  cgpa:number
  id:number
  firstName:string
  dob:string
  constructor(fN:string,cgpa:number,id:number,dob:string,srno:number){
    this.firstName=fN
    this.cgpa=cgpa
    this.id=id
    this.srno=srno
    this.dob=dob
  }
}
@Component({
  selector: 'app-looping',
  templateUrl: './looping.component.html',
  styleUrls: ['./looping.component.css']
})
export class LoopingComponent implements OnInit {
  Name!:string
  //now:Date=new Date()
 // time:number
  currentHour:string
  greet!:string
  show!:boolean
  constructor() {
    //this.time=this.now.getTime()
    var t=new Date()
    this.currentHour=t.toLocaleString('en-US', { hour: 'numeric', hour12: false })
    //this.currentHour="11 AM"
    console.log(Number(this.currentHour))  
    if (this.students.length ==0){
      this.show=true
    }else{
      this.show=false
    }
   }
  userOne:User=new User("Chinmay",21)
  userTwo:User=new User("Rahul",20)
  userThree:User=new User("Manish",22)
  userFour:User=new User("Raj",17)
  users:Array<User>=[
    this.userOne,this.userTwo,this.userThree,this.userFour
]
studentOne:Student=new Student("Chinmay",9,123,"16-02-2000",1)
studentTwo:Student=new Student("Manish",7.4,127,"18-01-2000",2)
studentThree:Student=new Student("Yash",5.9,128,"10-11-2000",3)
students:Array<Student>=[this.studentOne,this.studentTwo,this.studentThree]
//students:Array<Student>=[]
  ngOnInit(): void {
  }
  onkey(event:any):void{
    this.Name=event.target.value
    var Current_Hour=Number(this.currentHour)
    if( Current_Hour >= 8 && Current_Hour <= 11 ){
      this.greet="Good Morning !"
    } else if( Current_Hour >= 12 && Current_Hour <= 15) {
      this.greet="Good Afternoon ! "
    } else if (Current_Hour >= 16 && Current_Hour <= 20 ){
      this.greet="Good Evening !"
    } else {
      this.greet="Good night !"
    }
  }
  getBackgroundColor(cgpa:number):string{
    if (cgpa>=7.5){
      return 'green'
    }
    else if (cgpa>=6 && cgpa<=7.4){
      return 'yellow'
    }
    return 'red'
  }
}

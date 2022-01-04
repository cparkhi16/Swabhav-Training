import { Component, OnInit } from '@angular/core';


class User{
  firstName:string
  age:number
  constructor(fN:string,a:number){
    this.firstName=fN
    this.age=a
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
  constructor() {
    //this.time=this.now.getTime()
    var t=new Date()
    this.currentHour=t.toLocaleString('en-US', { hour: 'numeric', hour12: false })
    //this.currentHour="11 AM"
    console.log(Number(this.currentHour))  
   }
  userOne:User=new User("Chinmay",21)
  userTwo:User=new User("Rahul",20)
  userThree:User=new User("Manish",22)
  userFour:User=new User("Raj",17)
  users:Array<User>=[
    this.userOne,this.userTwo,this.userThree,this.userFour
]
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
}

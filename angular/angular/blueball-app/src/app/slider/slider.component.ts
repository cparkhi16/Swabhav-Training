import { Component, DoCheck, OnInit } from '@angular/core';

class Ball{
  id:number
  desc:number
  color:string
  constructor(id:number,desc:number){
    this.id=id
    this.desc=desc
    this.color="lightblue"
  }
}
@Component({
  selector: 'app-slider',
  templateUrl: './slider.component.html',
  styleUrls: ['./slider.component.css']
})

export class SliderComponent implements OnInit ,DoCheck{
   attempts:number=0
   allowedAttempts:number=0
   val!:number
   show:boolean=true
   balls:Array<Ball>=[]
   winNumber!:number
   ballColor!:string
   won:boolean=false
   remainingAttempts!:number
  constructor() { }

  ngOnInit(): void {
  }
  getRandomInt(max:number) {
    //this.attempts=max
    return Math.floor(Math.random() * max);
  }
  setTotalAttempts(val:number){
    console.log("Set total attempt called ",val)
    while (val) {
    this.allowedAttempts += val % 10;
    val = Math.floor(val / 10);
    }
    this.allowedAttempts=this.allowedAttempts+2
    console.log("Total allowed attempts ",this.allowedAttempts)
    this.remainingAttempts=this.allowedAttempts
  }
  getVal():void{
    let i:number
    this.winNumber=this.getRandomInt(this.val)
    console.log("Random number ",this.winNumber)
    if(this.winNumber==0){
      this.winNumber=this.getRandomInt(this.val)
    }
    this.show=false
    for(i=1;i<=this.val;i++){
    var b = new Ball(i,i)
    this.balls.push(b)
    }
    if(this.val<10){
    this.setTotalAttempts(0)
    }else{
      this.setTotalAttempts(this.val)
    }
  }
  ngDoCheck(): void {
      console.log(" Balls ",this.balls)
      // if(this.show==true){
      //   this.balls=[]
      // }
      for(let ball of this.balls){
        if(ball.color=='blue'){
          this.show=true
          //this.balls=[]
        }
      }
      
      if(this.show==true){
          this.balls=[]   
          this.allowedAttempts=0
          this.attempts=0
          this.winNumber=0
          //this.val=0  
      }
  }
  ballHit(ball:Ball):void{
    this.attempts=this.attempts+1
    this.remainingAttempts=this.allowedAttempts-this.attempts
    if(this.remainingAttempts<0){
      let msg="You lost all attempts"+" Correct number was "+this.winNumber
      alert(msg)
      this.show=true
    }
    else{
    console.log("ID Of ball hit ",ball.id)
    if(ball.id<this.winNumber){
      ball.color='green'
    }else if(ball.id>this.winNumber){
      ball.color='red'
    }else{
      ball.color='blue'
      alert("You Won !!")
      //this.won=true
     //this.show=true
     // alert("You Won !!")
      //this.balls=[]
    //   setTimeout(()=>{
    //     alert("You Won !!")
    //     this.show=true
    // },2000)
    }
  // }else{
  //   let msg="You lost all attempts"+" Correct number was "+this.winNumber
  //   alert(msg)
  //   this.show=true
  // }
  }
}
 
  getBallColor(ball:any):string{
    if(ball.color=='blue'){
      this.balls=[]
      //this.show=true
    }
    return ball.color
  }
}

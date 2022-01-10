import { Component, DoCheck, OnInit } from '@angular/core';

class Ball{
  id:number
  desc:number
  color:string
  isClicked:boolean
  constructor(id:number,desc:number){
    this.id=id
    this.desc=desc
    this.color="lightblue"
    this.isClicked=false
  }
}
@Component({
  selector: 'app-slider',
  templateUrl: './slider.component.html',
  styleUrls: ['./slider.component.css']
})

export class SliderComponent implements OnInit ,DoCheck{
  msg!:string
   gameResult!:boolean 
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
    // while (val) {
    // this.allowedAttempts += val % 10;
    // val = Math.floor(val / 10);
    // }
    // this.allowedAttempts=this.allowedAttempts+2
    // if(val!==10 && val!=11){
    // this.allowedAttempts=Math.round(Math.log2(val-10))
    // }else{
    //   this.allowedAttempts=1
    // }
    this.allowedAttempts=Math.round(Math.log2(val))
    console.log("Total allowed attempts ",this.allowedAttempts)
    this.remainingAttempts=this.allowedAttempts
  }
  getVal():void{
    let i:number
    this.winNumber=this.getRandomInt(this.val)
    console.log("Random number ",this.winNumber)
    if(this.winNumber==0){
      this.winNumber=this.getRandomInt(this.val)
      console.log("-= Changed Random number =-",this.winNumber)
    }
    this.show=false
    for(i=1;i<=this.val;i++){
    var b = new Ball(i,i)
    this.balls.push(b)
    }
   this.setTotalAttempts(this.val) 
  }
  ngDoCheck(): void {
      console.log(" Balls ",this.balls) 
      //console.log("In do check")
      if(this.show==true){
          this.balls=[]   
          this.allowedAttempts=0
          this.attempts=0
          this.winNumber=0
          //this.val=0  
      }
      //console.log("end of do check")
  }
  exit():void{
    this.show=true
    this.balls=[]   
    this.allowedAttempts=0
    this.attempts=0
    this.winNumber=0
    this.gameResult=false
  }
  restart(){
    console.log("--- restart ")
    this.gameResult=false
    this.balls=[]
    this.remainingAttempts=this.allowedAttempts
    this.attempts=0
    let i: number
    for(i=1;i<=this.val;i++){
      var b = new Ball(i,i)
      this.balls.push(b)
      }
    this.winNumber=this.getRandomInt(this.val)
    console.log("In restart random no ",this.winNumber)
  }
  showOriginalPage(){
    // setTimeout(()=>{
    //   this.show=true
    // },1000)
    this.show=true
  }
  ballHit(ball:Ball):void{
    if (ball.isClicked==false){
    this.attempts=this.attempts+1
    this.remainingAttempts=this.allowedAttempts-this.attempts
    ball.isClicked=true
    }
    if(this.remainingAttempts<0){
      let msg="You lost all attempts"+" Correct number was "+this.winNumber
     // alert(msg)
      this.gameResult=true
      this.msg=msg
      this.attempts=0
      this.remainingAttempts=this.allowedAttempts
    }
    else{
    console.log("ID Of ball hit ",ball.id)
    if(ball.id<this.winNumber){
      ball.color='green'
      this.gameResult=false
    }else if(ball.id>this.winNumber){
      ball.color='red'
      this.gameResult=false
    }else{
      ball.color='blue'
     // console.log("Here after blue ")
      this.gameResult=true
      this.msg= "You won !!"
      //this.showOriginalPage()
    }
  }
}
 
  getBallColor(ball:any):string{
    // if(ball.color=='blue'){
    //   this.balls=[]
    // }
    return ball.color
  }
}

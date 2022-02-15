import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { ball } from '../model/ball';
import { BallServiceService } from '../service/ball-service.service';

@Component({
  selector: 'app-ball-page',
  templateUrl: './ball-page.component.html',
  styleUrls: ['./ball-page.component.css']
})
export class BallPageComponent implements OnInit {
  //@ViewChild('content') content: any;
  @ViewChild('myDiv') myDiv!: ElementRef<HTMLElement>;
  numberOfBalls!:number;
  balls:ball[]=[];
  correctAns!:number;
  gameStatus!:string;
  modalDisplayStyle:string="none";
  numberOfAttemptsLeft:number=0;

  constructor(private ballService:BallServiceService) { 
    this.numberOfBalls=this.ballService.getNumberOfBalls();
    for(let i=1;i<=this.numberOfBalls;i++){
      this.balls.push({id:i,description:"ball"+i,color:"pink",clicked:false})
    }
    this.numberOfAttemptsLeft=Math.round(Math.log2(this.numberOfBalls));// Minimum number of guessing = log2(Upper bound – lower bound + 1)
  }

  ngOnInit(): void {
    this.numberOfBalls=this.ballService.getNumberOfBalls();
    let a=Math.random()*this.numberOfBalls
    this.correctAns=Math.floor(a)+1;
    console.log("%c correct ans","background-color:pink",this.correctAns);
    console.log(this.balls);
  }

  alterColorAndStatus(givenball:ball){
    if(givenball.id! < this.correctAns){
      givenball.color="green";
      this.gameStatus="loser";
    }
    else if(givenball.id! > this.correctAns){
      givenball.color="red";
      this.gameStatus="loser";
    }
    else{
      givenball.color="blue";
      this.gameStatus="winner";
    }
  }

  checkBall(givenball:ball){
    // console.log(givenball);
    givenball.clicked=true;
    this.numberOfAttemptsLeft=this.numberOfAttemptsLeft-1;
    if(this.numberOfAttemptsLeft!=-1){
      this.alterColorAndStatus(givenball);
      if(this.numberOfAttemptsLeft==0){
        let el: HTMLElement = this.myDiv.nativeElement;
        el.click();
      }
      // console.log(this.gameStatus);
    }
  }

  restartGame(){
    this.balls=[];
    this.numberOfAttemptsLeft=Math.round(Math.log2(this.numberOfBalls));// Minimum number of guessing = log2(Upper bound – lower bound + 1)
    for(let i=1;i<=this.numberOfBalls;i++){
      this.balls.push({id:i,description:"ball"+i,color:"pink"})
    }
    let a=Math.random()*this.numberOfBalls
    this.correctAns=Math.floor(a)+1;
    console.log("%c correct ans","background-color:pink",this.correctAns);
    console.log(this.balls);
  }

}

import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class BallServiceService {
  numberOfBalls!:number;
  constructor() { }

  getNumberOfBalls(){
    return this.numberOfBalls;
  }

  setNumberOfBalls(n:number){
    this.numberOfBalls=n;
  }
}

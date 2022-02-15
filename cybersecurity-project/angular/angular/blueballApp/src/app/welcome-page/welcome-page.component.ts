import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BallServiceService } from '../service/ball-service.service';

@Component({
  selector: 'app-welcome-page',
  templateUrl: './welcome-page.component.html',
  styleUrls: ['./welcome-page.component.css']
})
export class WelcomePageComponent implements OnInit {
  numberOfBalls:number=10;
  constructor(private router:Router,private ballService:BallServiceService) { }

  ngOnInit(): void {
  }

  startGame(){
    this.router.navigateByUrl('/play');
    this.ballService.setNumberOfBalls(this.numberOfBalls);
  }

}

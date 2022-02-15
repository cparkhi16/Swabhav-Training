import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-clock',
  templateUrl: './clock.component.html',
  styleUrls: ['./clock.component.css']
})
export class ClockComponent implements OnInit {
  date: any = Date.now();
  isHidden:boolean;
  buttonText:string;
  constructor() {
    this.isHidden=true;
    this.buttonText="show";
  }

  toggleClock(){
    this.isHidden=!this.isHidden;
    if(this.isHidden){
      this.buttonText="show";
    }
    else{
      this.buttonText="hide";
    }
  }

  ngOnInit(): void {
    setInterval(()=>{
      this.date=Date.now();
    },1000);
  }

}

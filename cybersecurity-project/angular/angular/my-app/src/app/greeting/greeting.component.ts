import { emitDistinctChangesOnlyDefaultValue } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-greeting',
  templateUrl: './greeting.component.html',
  styleUrls: ['./greeting.component.css']
})
export class GreetingComponent implements OnInit {
  currentHours:any;
  username:string="";
  message:string="welcome";
  constructor() { 

  }

  ngOnInit(): void {
  }

  onGreet(){
    this.currentHours=new Date().getHours();
    if(this.currentHours>=6 && this.currentHours<=12){
      this.message="good morning,"+this.username;
    }
    else if(this.currentHours>12 && this.currentHours<=17){
      this.message="good afternoon,"+this.username;
    }
    else if(this.currentHours>17 && this.currentHours<22){
      this.message="good evening,"+this.username;
    }
    else{
      this.message="good night,"+this.username;
    }
    console.log(this.username,this.currentHours);
  }

}

import { Component } from '@angular/core';

@Component({
  selector: 'custom-root',
  templateUrl: './customcomponent.component.html',
  styleUrls: ['./customcomponent.component.css']
})
export class CustomComponent {
  name!:string
  even:boolean;
  clickMessage!:string
  constructor(){
    this.name="CHINMAY PARKHI"
    console.log("Instance created ....")
    this.even=false;
  }
  onClickMe() {
    this.clickMessage = 'You are my hero!';
  }
  onKey(event: any) {
    console.log("New val",event.target.value)
    console.log("Type of val ",typeof(Number(event.target.value)))
    if (Number(event.target.value)%2==0) 
    this.even=true
    else
    this.even=false
  }
}

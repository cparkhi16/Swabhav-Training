import { Component,OnInit } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.css']
})
export class WelcomeComponent implements OnInit{
  title = 'welcome-app-v2';
  welcome!:string
  name:string="chInmay"
  today: number = Date.now();
  a:number=15
  pi: number = 3.14159265359;
  dateTime!:any
  constructor(){
    this.welcome="Welcome to angular"
  }
  ngOnInit(): void {
      this.welcome="Onit welcome"
      this.displayDateTime();
   setInterval(() => {
    this.displayDateTime(); 
  }, 1000);
  }
  displayDateTime() :void{
    this.dateTime = new Date()
   }
}

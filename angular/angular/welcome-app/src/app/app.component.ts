import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'welcome-app';
  name!:string
  constructor(){
    this.name="Good Morning"
    console.log("Instance created ....")
  }
}

import { Component,OnInit } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'my-app';
  name!:string;
  a: number = 0.568;
  data:number=0;
  count:number=0;
  constructor(){
    this.name="yogesh";
    this.count=0;
    console.log("in constructor");
  }

  clickedMe(){
    this.name="xyz";
    this.count=this.count+1;
  }

  changeFromParent(){
    this.data += 1;
  }

  ngOnInit(){
    console.log("in ngoninit");
  }
}

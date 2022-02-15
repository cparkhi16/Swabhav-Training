import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-test',
  templateUrl: './test.component.html',
  styleUrls: ['./test.component.css']
})
export class TestComponent implements OnInit {
    num!:any;
    isEven:boolean;
    text:string;
    type:string;
  constructor() {
    this.num=0;
    this.isEven=true;
    this.text="";
    this.type="";
  }

  check(event: any) {
    this.num= event.target.value;
    if(this.num%2==0){
        this.isEven=true;
        this.type="Even";
    }
    else{
        this.isEven=false;
        this.type="Odd";
    }
};

  ngOnInit(): void {
  }

}

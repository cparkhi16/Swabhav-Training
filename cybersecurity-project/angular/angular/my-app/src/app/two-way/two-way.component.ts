import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-two-way',
  templateUrl: './two-way.component.html',
  styleUrls: ['./two-way.component.css']
})
export class TwoWayComponent implements OnInit {
  age:number=0;
  lang!:string;
  constructor() {
    this.age=103;
   }

  ngOnInit(): void {
  }

}
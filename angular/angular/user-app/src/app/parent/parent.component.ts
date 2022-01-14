import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-parent',
  templateUrl: './parent.component.html',
  styleUrls: ['./parent.component.css']
})
export class ParentComponent implements OnInit {
  rating!:number
  data!:any
  constructor() { }

  ngOnInit(): void {
  }
  parentHandler(e:any):void{
    console.log("Parent handler called !")
    this.rating=e
  }
}

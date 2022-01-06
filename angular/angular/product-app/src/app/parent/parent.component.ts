import { Component, DoCheck, OnChanges, OnDestroy, OnInit, SimpleChanges } from '@angular/core';

@Component({
  selector: 'app-parent',
  templateUrl: './parent.component.html',
  styleUrls: ['./parent.component.css']
})
export class ParentComponent implements OnInit,OnChanges,OnDestroy ,DoCheck{
  itemOfParent:any=100
  constructor() { }

  ngOnInit(): void {
    console.log("Init of parent ")
  }
  ngOnChanges(changes: SimpleChanges): void {
      console.log("Changes in parent ",changes)
  }
  ngOnDestroy(): void {
    console.log("Parent destroyed !!")
  }
  ngDoCheck(): void {
    console.log(" --= PARENT =-- Called after ngOnChanges and first time after oninit--> ")
}
  parentHandler(e:any){
    console.log("Parent hadnler  ",e )
  }
}

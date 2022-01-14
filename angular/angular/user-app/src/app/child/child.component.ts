import { Component, DoCheck, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges } from '@angular/core';

@Component({
  selector: 'app-child',
  templateUrl: './child.component.html',
  styleUrls: ['./child.component.css']
})
export class ChildComponent implements OnInit,DoCheck{
  starRating:number=0;
  //currentRate = 3.14;
  @Output() childRating:EventEmitter<number>=new EventEmitter()
  @Input() rating: number=1;
  visible!:string
  constructor() { }

  ngOnInit(): void {
  }
 ngDoCheck(): void {
      console.log("Value here rating ",this.starRating)
      this.childRating.emit(this.starRating)
      //this.changesOnChild(this.starRating)
  }
  getOuterDivWidth():string{
    let w=this.rating*54.78
    this.visible="hidden"
    return w +'px'
  }
  getInnerDivWidth():string{
    return '400px'
  }
  
}

import { AfterContentChecked, AfterContentInit, AfterViewInit, Component, ContentChild, DoCheck, EventEmitter, Input, OnChanges, OnDestroy, OnInit, Output, SimpleChanges, ViewChild } from '@angular/core';

@Component({
  selector: 'app-child',
  templateUrl: './child.component.html',
  styleUrls: ['./child.component.css']
})
export class ChildComponent implements OnInit,OnChanges,OnDestroy,AfterContentInit,DoCheck,AfterViewInit {
  @Input() itemOfChild:any=0
  @Output() clickChild:EventEmitter<any>=new EventEmitter()
  constructor() { 
    console.log("Value of parent template ",this.valParent)
    console.log("Value of child template ",this.myVal)
  }
  @ContentChild('parentTemplate') valParent:any
  @ViewChild('myTemplate') myVal:any
  ngOnInit(): void {
    console.log("Init of parent ")
    console.log("Value of parent template ",this.valParent)
    console.log("Value of child template ",this.myVal)
  }
  ngOnChanges(changes: SimpleChanges): void {
      console.log("Changes in child ",changes)
  }
  ngOnDestroy(): void {
      console.log("Child destroyed !!")
  }
  ngAfterContentInit(): void {
      console.log("Within after content init")
      console.log("Value of parent template ",this.valParent)
      console.log("Value of child template ",this.myVal)
  }
  ngAfterViewInit(): void {
    console.log("Within after view init ==============")
    console.log("Value of parent template ",this.valParent)
    console.log("Value of child template ",this.myVal)
  }
  ngDoCheck(): void {
      console.log(" --= Child =-- Called after ngOnChanges and first time after oninit--> ")
  }
  handleClick():void{
    console.log("p clicked --------")
    this.clickChild.emit({id:123,name:"Chinmay"})
  }
}

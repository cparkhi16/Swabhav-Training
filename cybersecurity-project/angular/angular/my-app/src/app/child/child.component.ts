import { Component, ContentChild, EventEmitter, Input, OnChanges, OnDestroy, OnInit, Output, SimpleChanges, ViewChild } from '@angular/core';

@Component({
  selector: 'app-child',
  templateUrl: './child.component.html',
  styleUrls: ['./child.component.css']
})
export class ChildComponent implements OnInit,OnChanges,OnDestroy {
  @Input()itemFromChild:any;
  @Output() newItemEvent = new EventEmitter<any>();
  @ContentChild('template1') template1:any;
  @ViewChild('template2') template2:any;
  constructor() {
    console.log("%c child constructor","background-color:pink");
  }

  ngOnInit(): void {
    console.log("child ngOnInit");
  }

  ngOnChanges(changes: SimpleChanges): void {
      console.log("%c child ngOnChanges","background-color:pink",changes);
  }

  ngOnDestroy(){
    console.log("%c child ngOnDestroy","background-color:pink");
  }

  ngAfterContentInit(){
    console.log("%c child ngAfterContentInit","background-color:pink");
    console.log(this.template1);
    console.log(this.template2);
  }

  ngAfterContentChecked(){
    console.log("%c child ngAfterContentChecked","background-color:pink");
    console.log(this.template1);
    console.log(this.template2);
  }

  ngAfterViewInit(){
    console.log("%c child ngAfterViewInit","background-color:pink");
    console.log(this.template1);
    console.log(this.template2);
  }

  ngAfterViewChecked(){
    console.log("%c child ngAfterViewChecked","background-color:pink");
    console.log(this.template1);
    console.log(this.template2);
  }

  ngDoCheck(){
    console.log("%c child ngDoChcek","background-color:pink");
  }

  sendMsgToParent(){
    this.newItemEvent.emit({name:"hellomsg",age:23});
  }


}

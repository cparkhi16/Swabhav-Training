import { Component, OnChanges, OnInit, SimpleChanges } from '@angular/core';

@Component({
  selector: 'app-hooks',
  templateUrl: './hooks.component.html',
  styleUrls: ['./hooks.component.css']
})
export class HooksComponent implements OnInit,OnChanges {
  itemFromParent:number=10;
  flag=false;
  constructor() {
    console.log("%c parent constructor","background-color:yellow");
  }

  ngOnInit(): void {
    console.log("%c parent ngOnInit","background-color:yellow");
  }

  ngOnChanges(changes: SimpleChanges): void {
      console.log("%c  parent ngOnChanges","background-color:yellow",changes);
  }

  ngDoCheck(){
    console.log("%c parent ngDoChcek","background-color:yellow");
  }

  getmsg(event:any){
    console.log(event);
  }

}

import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-child',
  templateUrl: './child.component.html',
  styleUrls: ['./child.component.css']
})
export class ChildComponent implements OnInit {
  @Input() rating!:number;
  constructor() { }

  ngOnInit(): void {
  }

  getWidth(){
    return (this.rating*83.3)+'px';
  }

}

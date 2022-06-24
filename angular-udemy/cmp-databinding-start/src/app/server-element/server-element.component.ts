import { Component, Input, OnInit, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-server-element',
  templateUrl: './server-element.component.html',
  styleUrls: ['./server-element.component.css'],
  //encapsulation : ViewEncapsulation.Emulated // by default it is emulated that is view encapsulation is enabled ( css to component binding)
  //encapsulation : ViewEncapsulation.None // it disables view encapsulation
})
export class ServerElementComponent implements OnInit {
  @Input('srvElement') element : {type:string,name:string,content:string};
  constructor() { }

  ngOnInit(): void {
  }

}

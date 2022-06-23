import { Component, OnInit } from '@angular/core';

@Component({
  //selector: '[app-servers]',as attribute
  //selector: '.app-servers', //as class in CSS selector
  selector : 'app-servers',// as an elm selector
  //template: `<app-server></app-server> <h3>Style test</h3>`,
  templateUrl: './servers.component.html',
  styleUrls: ['./servers.component.css']
  // styles : [`
  // h3{
  //   color: red
  // }`]
})
export class ServersComponent implements OnInit {
  inputData = "chinmay"
  constructor() { }

  ngOnInit(): void {
  }
  onInput(event:Event){
    this.inputData= (<HTMLInputElement>event.target).value;
  }
  
}

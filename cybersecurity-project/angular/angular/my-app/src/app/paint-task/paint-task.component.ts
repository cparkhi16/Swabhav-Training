import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-paint-task',
  templateUrl: './paint-task.component.html',
  styleUrls: ['./paint-task.component.css']
})
export class PaintTaskComponent implements OnInit {
  data:string="";
  fontSize!:string;
  colors:string[]=["red","blue","green","pink","white","black","purple","indigo","grey","yellow","orange"];
  selectedBgColor:string="white";
  selectedFontColor:string="black";
  constructor() { }

  ngOnInit(): void {
  }

}

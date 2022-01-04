import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-two-way',
  templateUrl: './two-way.component.html',
  styleUrls: ['./two-way.component.css']
})
export class TwoWayComponent implements OnInit {
  name:string
  myName!:string
  font!:string
  SelectedColor!:string
  TextColor!:string
  language!:string
  colors:string[]=["blue","red","green"]
  constructor() { 
    this.name="Chinmay"
  }

  ngOnInit(): void {
  }
  onChange(){
    console.log("Text to be displayed ",this.myName)
    console.log("Font is ",this.font)
    console.log("Background color is ",this.SelectedColor)
    console.log("Selected Text color ",this.TextColor)
    if (this.SelectedColor==this.TextColor){
      alert("Please enter different font and background colors")
      location.reload()
    }
    }

}

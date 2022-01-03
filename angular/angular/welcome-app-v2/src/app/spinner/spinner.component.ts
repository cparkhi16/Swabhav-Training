import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-spinner',
  templateUrl: './spinner.component.html',
  styleUrls: ['./spinner.component.css']
})
export class SpinnerComponent implements OnInit {
  spinnerURL:string="../../assets/loading-gif.gif"
  showSpinner:boolean
  visibility:string=""
  btnval!:string
  constructor() {
    this.btnval="hide"
    this.showSpinner=true
   }
  
  ngOnInit(): void {
  }
  onButtonClick():void{
    if(this.showSpinner==true){
    this.btnval="show"
    this.visibility="hidden"
    this.showSpinner=false
    }
    else if (this.showSpinner==false){
      this.btnval="hide"
      this.visibility=""
      this.showSpinner=true
    }
  }
}

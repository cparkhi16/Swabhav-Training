import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-spinner',
  templateUrl: './spinner.component.html',
  styleUrls: ['./spinner.component.css']
})
export class SpinnerComponent implements OnInit {
  showSpinner:boolean=false;
  imageSrc:string[]=["img1.jpg","img2.jfif","img3.jfif","img4.jfif"];
  currentSrc:string="img1.jpg";
  spinnerSrc:string="../../assets/Winter.gif";
  count:number=0;
  isSpinnerHidden:string="hidden";
  buttonName:string="show";
  constructor() { }

  toggleSpinner(){
    if(this.isSpinnerHidden=="hidden"){
      this.isSpinnerHidden="";
      this.buttonName="hide";
    }else{
      this.isSpinnerHidden="hidden";
      this.buttonName="show";
    }
  }

  getUrl()
{
  return "url('../../assets/"+this.currentSrc+"')";
}

  ngOnInit(): void {
    setInterval(()=>{
      this.currentSrc=this.imageSrc[this.count];
      if(this.count==this.imageSrc.length-1){
        this.count=0;
      }
      else{
        this.count=this.count+1;
      }
    },5000)
  }
}

import { Component,OnInit,Pipe, PipeTransform } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.css']
})
export class WelcomeComponent implements OnInit{
  randomNumber:number=0
  imgURL:string="../assets/scenery.jfif"
  imageList=["../assets/scenery.jfif","../assets/peacock.jfif","../assets/tiger.jfif"]
  title = 'welcome-app-v2';
  welcome!:string
  name:string="chInmay"
  today: number = Date.now();
  a:number=15
  pi: number = 3.14159265359;
  dateTime!:any
  show:boolean=true;
  file = { name: 'logo.svg', size: 2120109, type: 'image/svg' };
  constructor(){
    this.welcome="Welcome to angular"
  }
  ngOnInit(): void {
      this.welcome="Onit welcome"
      this.displayDateTime();
   setInterval(() => {
    this.displayDateTime(); 
  }, 1000);
  setInterval(() => {
    this.displayImage(); 
  }, 1000);
  }
  displayImage():void{
    this.imgURL=this.imageList[this.randomNumber]
    this.randomNumber=this.randomNumber+1
    if (this.randomNumber==this.imageList.length){
      this.randomNumber=0
    }
  }
  displayDateTime() :void{
    if (this.show==true)
    this.dateTime = new Date()
    else
    this.dateTime=""
   }
   toggleClock(){
    this.show=!this.show
  }
}
@Pipe({ name: 'filesize' })
export class FileSizePipe implements PipeTransform {
  transform(size: number,extension:string): string {
    return (size / (1024 * 1024)).toFixed(2) + extension;
  }
}
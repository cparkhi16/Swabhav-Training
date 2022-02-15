import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-parent',
  templateUrl: './parent.component.html',
  styleUrls: ['./parent.component.css']
})
export class ParentComponent implements OnInit {
  rating!:number;
  ratingFromChild!:number;
  constructor() { }

  ngOnInit(): void {
  }

  validateRating():boolean{
    if(this.rating===undefined || this.rating<0 || this.rating>5){
      return true;
    }
    return false;
  }

  showRating(event:any){
    this.ratingFromChild=event;
  }

}

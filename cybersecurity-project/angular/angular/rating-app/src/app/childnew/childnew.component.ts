import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import {
  trigger,
  state,
  style,
  animate,
  transition,
} from '@angular/animations';

@Component({
  selector: 'app-childnew',
  templateUrl: './childnew.component.html',
  styleUrls: ['./childnew.component.css'],
  animations: [
    // Each unique animation requires its own trigger. The first argument of the trigger function is the name
    trigger('rotatedState', [
      state('default', style({ transform: 'rotate(0)' })),
      state('rotated', style({ transform: 'rotate(-360deg)' })),
      transition('rotated => default', animate('800ms ease-out')),
      transition('default => rotated', animate('800ms ease-in')),
    ]),
  ],
})
export class ChildnewComponent implements OnInit {
  @Output() ratingDone=new EventEmitter<number>();
  //rstate: string = 'default';
  stars:any[]=[
    {id:1,color:"grey",state:"default"},
    {id:2,color:"grey",state:"default"},
    {id:3,color:"grey",state:"default"},
    {id:4,color:"grey",state:"default"},
    {id:5,color:"grey",state:"default"},
  ];
  isStarClicked=false;
  constructor() { }

  ngOnInit(): void {
  }

  starClicked(id:number){
    for(let i=0;i<id;i++){
      this.stars[i].color="yellow";
    }
    this.isStarClicked=true;
    this.ratingDone.emit(id);
  }

  starHovered(id:number){
    if(!this.isStarClicked){
      for(let i=0;i<id;i++){
        this.stars[i].color="yellow";
      }
    }
    this.stars[id-1].state = this.stars[id-1].state === 'default' ? 'rotated' : 'default';
  }

  starOut(id:number){
    if(!this.isStarClicked){
      this.stars[id-1].color="grey";
    }
  }

  reset(){
    for(let i=0;i<this.stars.length;i++){
      this.stars[i].color="grey";
    }
    this.isStarClicked=false;
  }

}

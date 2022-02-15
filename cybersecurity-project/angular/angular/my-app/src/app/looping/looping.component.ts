import { Component, Input, OnInit, SimpleChanges } from '@angular/core';

class User{
  firstName:string
  age:number
  constructor(fN:string,a:number){
    this.firstName=fN
    this.age=a
  }
}

@Component({
  selector: 'app-looping',
  templateUrl: './looping.component.html',
  styleUrls: ['./looping.component.css']
})


export class LoopingComponent implements OnInit {
  @Input() parentData!:any;
  users:any[]=[
    {fname:"max",lname:"wek",age:12},
    {fname:"mina",lname:"tyao",age:34},
    {fname:"wan",lname:"qik",age:3},
    {fname:"eric",lname:"desnaer",age:32},
    {fname:"brenda",lname:"uok",age:15},
    {fname:"kevin",lname:"tyao",age:89},
    {fname:"turwi",lname:"wrka",age:10},
  ];
  heroes:any[]=[
    {name:"x",emotion:"happy"},
    {name:"y",emotion:"sad"},
    {name:"z",emotion:"filmy"},
    {name:"a",emotion:"zero"}
  ]
  changeLog: any;
  constructor() { 
    var user1:User=new User("x",2);
    console.log(user1);
  }

  addSomething(){
    this.users.push({fname:"something",lname:"something",age:45});
  }

  changeFromChild(){
    this.parentData -= 1;
  }

  ngOnInit(): void {
  }

  ngOnChanges(changes: SimpleChanges) {
    console.log(changes)
  }
  ngDoCheck(){
    console.log("DO CHECK")
  }

}

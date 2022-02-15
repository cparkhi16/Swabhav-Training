import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-student',
  templateUrl: './student.component.html',
  styleUrls: ['./student.component.css']
})
export class StudentComponent implements OnInit {
  students:any[]=[];
  // students:any[]=[
  //   {id:"23",name:"martin",cgpa:9.8,dob:new Date(1985,11,21)},
  //   {id:"89",name:"kevin",cgpa:3.2,dob:new Date(1966,1,13)},
  //   {id:"93",name:"brenda",cgpa:6.6,dob:new Date(1999,5,4)},
  //   {id:"45",name:"austin",cgpa:7.9,dob:new Date(2000,6,3)},
  //   {id:"10",name:"kelly",cgpa:7.2,dob:new Date(1977,15,10)}
  // ]
  constructor() { }

  getBgColor(cgpa:number):string{
    if(cgpa>=7.5){
      return 'green';
    }else if(cgpa>6 && cgpa<7.5){
      return 'yellow';
    }
    return 'red';
  }

  ngOnInit(): void {
  }

}

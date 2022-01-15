import { ObsService } from './../myservice/obs.service';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-course',
  templateUrl: './course.component.html',
  styleUrls: ['./course.component.css']
})
export class CourseComponent implements OnInit {
  courses:any[]=[]
  isCourseData:boolean=false
  courseName!:any
  display="none"
  constructor(private obs:ObsService) { 
    console.log("Course component constructor called !!")
  }

  ngOnInit(): void {
    this.updateCourseView()
  }
  updateCourseView(){
    this.obs.getAllCourses().subscribe({
      next:(data:any)=>{
        console.log(" DATA COURSE FROM API ",data)
        if(data.length!=0){
          this.isCourseData=true
          this.courses=data
      }
      }
    }
    )
  }
  deleteCourse(course:any){
    console.log("Course delete ",course.ID)
  }
  addCourse(){
    console.log("New course name ",this.courseName)
    this.obs.addCourse(this.courseName).subscribe({
      next:(data)=>{
        this.updateCourseView()
      },
      error:(err)=>{
        console.log("Error adding new course ",err)
      }
    })
    this.display="none"
  }
  onCloseHandled(){
    this.display="none"
  }
  openModal(){
    this.display="block"
  }
}

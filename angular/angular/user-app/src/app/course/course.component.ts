import { ObsService } from './../myservice/obs.service';
import { Component, OnInit } from '@angular/core';
import { FormGroup,Validators,FormControl } from '@angular/forms';
@Component({
  selector: 'app-course',
  templateUrl: './course.component.html',
  styleUrls: ['./course.component.css']
})
export class CourseComponent implements OnInit {
  myGroup:any
  courses:any[]=[]
  isCourseData:boolean=false
  courseName!:any
  display="none"
  displayCourseModal="none"
  courseToBeUpdated:any
  constructor(private obs:ObsService) { 
    console.log("Course component constructor called !!")
  }

  ngOnInit(): void {
    this.updateCourseView()
    this.myGroup = new FormGroup({
      updatedCourseName : new FormControl('',[Validators.required,Validators.maxLength(50)]),
    })
  }
  openUpdateCourseModal(course:any){
    this.displayCourseModal="block"
    this.courseToBeUpdated=course
    console.log("Course updated id ",course.ID)
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
    this.obs.deleteCourse(course.ID).subscribe({
      next:(data)=>{
        this.updateCourseView()
      },
      error:(err)=>{
        console.log("Error deleting course ",err)
      }
    })
  }
  updateCourse(myGroup:any){
    this.displayCourseModal="none"
    console.log("Updating course ",this.courseToBeUpdated.ID)
    console.log("New course name ",myGroup.value.updatedCourseName)
    this.obs.updateCourse(this.courseToBeUpdated.ID,myGroup.value.updatedCourseName).subscribe({
      next:(data)=>{
        this.updateCourseView()
      },
      error:(err)=>{
        console.log("Error updating course name ",err)
      }
    })
  }
  closeCourseModal(){
    this.displayCourseModal="none"
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

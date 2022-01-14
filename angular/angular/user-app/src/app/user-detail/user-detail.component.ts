import { Component, Input, OnInit } from '@angular/core';
import { ObsService } from '../myservice/obs.service';
import { FormControl, FormGroup, Validators,ValidationErrors,AbstractControl } from '@angular/forms';
@Component({
  selector: 'app-user-detail',
  templateUrl: './user-detail.component.html',
  styleUrls: ['./user-detail.component.css']
})
export class UserDetailComponent implements OnInit {
  hobbies:any[]=[]
  courses:any[]=[]
  passport:any
  isCourseData!:boolean
  isPassportData!:boolean
  isHobbyData!:boolean
  constructor(private obs:ObsService) { }
  @Input() userID:any
  ngOnInit(): void {
    this.obs.getHobbiesForUser(this.userID).subscribe({
      next:(data:any)=>{
      this.hobbies=data
      if(this.hobbies.length!=0){
      this.isHobbyData=true
      }
      console.log("Hobbies data -",data)}
    })
    this.obs.getCourseAndPassportForUser(this.userID).subscribe({
      next:(data:any)=>{
        this.courses=data.Courses
        this.passport=data.Passport
        if (this.courses.length!=0)
        {
        this.isCourseData=true
        }
        if(this.passport.PassportID!=0){
        this.isPassportData=true
        }
        console.log("Is passport data ",this.isPassportData)
        console.log("Is courses data ",this.isCourseData)
        console.log("Course data -",data.Courses)
        console.log("Passport data",data.Passport)
      }
    })
  }
  deleteHobby(hobby:any){
    console.log("HobbyID ",hobby)
    //this.hobbies.splice()
    this.obs.deleteHobbyForUser(hobby.ID).subscribe((data)=>{
      console.log("Data ",data)
    })
  }
  deleteEnrolledCourse(course:any){
    console.log("Delete course for user ",course)
  }
  deleteEnrolledPassport(passport:any){
    console.log("Delete passport for user ",passport)
  }

}

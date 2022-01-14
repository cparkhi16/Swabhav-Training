import { Component, Input, OnInit } from '@angular/core';
import { ObsService } from '../myservice/obs.service';
import { FormControl, FormGroup, Validators,ValidationErrors,AbstractControl } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-user-detail',
  templateUrl: './user-detail.component.html',
  styleUrls: ['./user-detail.component.css']
})
export class UserDetailComponent implements OnInit {
  // userForm:any
  newPassportID!:number
  newExpiryDateForPassport!:any
  userID:any
  hobbyName!:string
  coursesInDB:any[]=[]
  userCourses:any=[]
  hobbies:any[]=[]
  courses:any[]=[]
  passport:any
  SelectedCourse:string=""
  isCourseData:boolean=false
  isPassportData:boolean=false
  isHobbyData:boolean=false
  display = "none";
  displayPassport="none";
  constructor(private obs:ObsService,private router:Router,
    private activatedRoute:ActivatedRoute) { 
      this.activatedRoute.paramMap.subscribe(params=>{
        this.userID=params.get('userId') //+ string to number
      })
    }
 // @Input() userID:any
  
  ngOnInit(): void {
    // this.userForm = new FormGroup({
    //   hobbyName: new FormControl('',[Validators.required,Validators.maxLength(30)])
    // })
    this.obs.getHobbiesForUser(this.userID).subscribe({
      next:(data:any)=>{
      this.hobbies=data
      if(this.hobbies.length!=0){
      this.isHobbyData=true
      }
      console.log("Hobbies data -",data)}
    })
    this.getAllDBCourses()
    this.obs.getCourseAndPassportForUser(this.userID).subscribe({
      next:(data:any)=>{
        this.courses=data.Courses //user enrolled courses
        this.passport=data.Passport
        if (this.courses.length!=0)
        {
        this.isCourseData=true
        // this.obs.getAllCourses().subscribe({
        //   next:(data:any)=>{
            for(let i=0;i<this.coursesInDB.length;i++){
              console.log("Courses ",this.coursesInDB[i].Name)
              for(let j=0;j<this.courses.length;j++){
                if(this.coursesInDB[i].Name!=this.courses[j].Name){
                  this.userCourses.push(this.coursesInDB[i].Name)
                  //this.coursesInDB.push(data[i])
                }
        //       }
        //     }
        //   }
        // })
        }
      }
    }
      else{
          // this.obs.getAllCourses().subscribe({
          //   next:(data:any)=>{
          //     console.log(" DATA COURSE FROM API ",data)
              for(let i=0;i<this.coursesInDB.length;i++){
                console.log("Courses ",this.coursesInDB[i].Name)
                this.userCourses.push(this.coursesInDB[i].Name) //all the courses in db in which user can enroll
                //this.coursesInDB.push(data[i])
              }
              
          //   }
          // })
        }
        if(this.passport.PassportID!=0){
        this.isPassportData=true
        }
        console.log("Is passport data ",this.isPassportData)
        console.log("Is courses data ",this.isCourseData)
        console.log("Course data -",data.Courses)
        console.log("Passport data",data.Passport)
      },
      error:(err)=>{
        console.log("Error in getting course or passport data ",err)
      }
    })
  }
  getAllDBCourses(){
    this.obs.getAllCourses().subscribe({
      next:(data:any)=>{
        console.log(" DATA COURSE FROM API ",data)
        for(let i=0;i<data.length;i++){
          this.coursesInDB.push(data[i])
        }
      }
    }
    )
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
  openModal() {
    this.display = "block";
  }
  openModalForPassport(){
    this.displayPassport="block";
  }
  onCloseHandledForPassport(){
    this.displayPassport="none";
    if(this.isCourseData==false){
      console.log("Need to call api ",this.newPassportID,this.newExpiryDateForPassport)
      this.obs.addPassportDetailsForUser(this.userID,this.newPassportID,this.newExpiryDateForPassport).subscribe({
        error:(err)=>{
          console.log("Error in add passport detail ",err)
        }
      })
    }
  }
  onCloseHandled() {
    console.log("New hobby name ",this.hobbyName)
    this.obs.addUserHobby(this.userID,this.hobbyName).subscribe({
      error:(err)=>{
        console.log("Error in new add hobby ",err)
      }
    })
    this.display = "none";
  }
  onCourseChange(){
    console.log("Course add --- ",this.SelectedCourse)
    for(let c of this.coursesInDB){
      if (c.Name==this.SelectedCourse){
        console.log("ID of selected course to be added ",c.ID)
        this.obs.enrollUserCourse(this.userID,c.ID).subscribe({
          error:(err)=>{
            console.log("Error adding course ",err)
          }
        })
        alert("Course added !!")
      }
    }
  }
}

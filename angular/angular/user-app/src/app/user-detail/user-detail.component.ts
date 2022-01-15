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
  date:any
  newPassportID!:any
  newExpiryDateForPassport!:any
  userID:any
  hobbyName:string=""
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
  updateView(){
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
        this.courses=data.Courses //user enrolled courses
        this.passport=data.Passport
        if (this.courses.length!=0)
        {
        this.isCourseData=true
        this.userCourses=this.coursesInDB
            for(let i=0;i<this.courses.length;i++){
              for(let j=0;j<this.userCourses.length;j++){
               // console.log("Courses name in db ",this.coursesInDB[i].Name)
                // if(this.coursesInDB[i].Name!==this.courses[j].Name){
                //   console.log("Courses name of user enrolled  in db ",this.courses[j].Name)
                //   this.userCourses.push(this.coursesInDB[i].Name)
                // }
                if(this.userCourses[j].Name===this.courses[i].Name){
                  console.log("same course in to be enrolled courses found ")
                  this.userCourses.splice(j,1)
                }
        }
        console.log("User courses ",this.userCourses)
      }
    }
      else{
              for(let i=0;i<this.coursesInDB.length;i++){
                console.log("Courses ",this.coursesInDB[i].Name)
                this.userCourses.push(this.coursesInDB[i]) //all the courses in db in which user can enroll
              }
        }
        if(this.passport.PassportID!=0){
        this.isPassportData=true
        }else{
          this.isPassportData=false
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
  ngOnInit(): void {
    // this.userForm = new FormGroup({
    //   hobbyName: new FormControl('',[Validators.required,Validators.maxLength(30)])
    // })
    // this.obs.getHobbiesForUser(this.userID).subscribe({
    //   next:(data:any)=>{
    //   this.hobbies=data
    //   if(this.hobbies.length!=0){
    //   this.isHobbyData=true
    //   }
    //   console.log("Hobbies data -",data)}
    // })
    this.getAllDBCourses()
    this.updateView()
    // this.obs.getCourseAndPassportForUser(this.userID).subscribe({
    //   next:(data:any)=>{
    //     this.courses=data.Courses //user enrolled courses
    //     this.passport=data.Passport
    //     if (this.courses.length!=0)
    //     {
    //     this.isCourseData=true
    //     // this.obs.getAllCourses().subscribe({
    //     //   next:(data:any)=>{
    //         for(let i=0;i<this.coursesInDB.length;i++){
    //           console.log("Courses ",this.coursesInDB[i].Name)
    //           for(let j=0;j<this.courses.length;j++){
    //             if(this.coursesInDB[i].Name!=this.courses[j].Name){
    //               this.userCourses.push(this.coursesInDB[i].Name)
    //               //this.coursesInDB.push(data[i])
    //             }
    //     //       }
    //     //     }
    //     //   }
    //     // })
    //     }
    //   }
    // }
    //   else{
    //       // this.obs.getAllCourses().subscribe({
    //       //   next:(data:any)=>{
    //       //     console.log(" DATA COURSE FROM API ",data)
    //           for(let i=0;i<this.coursesInDB.length;i++){
    //             console.log("Courses ",this.coursesInDB[i].Name)
    //             this.userCourses.push(this.coursesInDB[i].Name) //all the courses in db in which user can enroll
    //             //this.coursesInDB.push(data[i])
    //           }
              
    //       //   }
    //       // })
    //     }
    //     if(this.passport.PassportID!=0){
    //     this.isPassportData=true
    //     }
    //     console.log("Is passport data ",this.isPassportData)
    //     console.log("Is courses data ",this.isCourseData)
    //     console.log("Course data -",data.Courses)
    //     console.log("Passport data",data.Passport)
    //   },
    //   error:(err)=>{
    //     console.log("Error in getting course or passport data ",err)
    //   }
    // })
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
      this.updateView()
    })
  }
  deleteEnrolledCourse(course:any){
    console.log("Delete course for user ",course)
    this.obs.deleteCourseForUser(this.userID,course.ID).subscribe({
      next:(data)=>{
        this.updateView()
        this.userCourses.push(course)
      },
      error:(err)=>{
        console.log("Error unenrolling course for user ",err)
      }
    })
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
    if(this.isPassportData==false){
      console.log("Need to call api ",this.newPassportID,this.newExpiryDateForPassport.day+"-"+this.newExpiryDateForPassport.year)
      this.newExpiryDateForPassport=this.newExpiryDateForPassport.year+"-"+this.newExpiryDateForPassport.month+"-"+this.newExpiryDateForPassport.day
      console.log("Need to call api  -- ",this.newPassportID,this.newExpiryDateForPassport)
      this.obs.addPassportDetailsForUser(this.userID,this.newPassportID,this.newExpiryDateForPassport).subscribe({
        next:(data)=>{
          this.updateView()
        },
        error:(err)=>{
          console.log("Error in add passport detail ",err)
        }
      })
    }
  }
  onCloseHandled() {
    console.log("New hobby name ",this.hobbyName)
    if(this.hobbyName!=""){
    this.obs.addUserHobby(this.userID,this.hobbyName).subscribe({
      next:(data)=>{
        this.updateView()
      },
      error:(err)=>{
        console.log("Error in new add hobby ",err)
      }
    })
  }
    this.display = "none";
  }
  onCourseChange(event:any){
    this.SelectedCourse=event.target.value
    console.log("Course add --- ",this.SelectedCourse)
    if(this.SelectedCourse!=="default"){
    for(let c of this.coursesInDB){
      if (c.Name==this.SelectedCourse){
        console.log("ID of selected course to be added ",c.ID)
        this.obs.enrollUserCourse(this.userID,c.ID).subscribe({
          next:(data)=>{
          this.updateView()
        },
          error:(err)=>{
            console.log("Error adding course ",err)
          }
        })
        alert("Course added !!")
      }
    }
  }
  }
}

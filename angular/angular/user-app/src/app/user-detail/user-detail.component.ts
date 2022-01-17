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
  userName!:string
  myGroup:any
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
  displayUpdatePassportModel="none";
  updateExpiryDateForPassport:any
  constructor(private obs:ObsService,private router:Router,
    private activatedRoute:ActivatedRoute) { 
      this.activatedRoute.paramMap.subscribe(params=>{
        this.userID=params.get('userId') //+ string to number
      })
    }
 goToProfile(){
  this.router.navigate(["profile/",this.userID])
 }
 goToCourse(){
   this.router.navigate(["courses"])
 }
 openUpdatePassportModel(passport:any){
   console.log("Modal opened to update passport ",passport.ID)
   this.displayUpdatePassportModel="block"
 }
 updatePassport(myGroup:any){
  let updatedExpiryDate=""
  this.displayUpdatePassportModel="none"
  if(myGroup.value.updateExpiryDateForPassport!=""){
  updatedExpiryDate=myGroup.value.updateExpiryDateForPassport.year+"-"+myGroup.value.updateExpiryDateForPassport.month+"-"+myGroup.value.updateExpiryDateForPassport.day
  }
   console.log("Passport ID to be updated ",myGroup.value.updatedPassportID,updatedExpiryDate,this.passport.ID)
   this.obs.updatePassport(this.passport.ID,myGroup.value.updatedPassportID,updatedExpiryDate).subscribe({
     next:(data)=>{
       this.updateView()
       this.displayPassport="none"
     },
     error:(err)=>{
      console.log("Error updating passport ",err)
     }
   })
 }
 closeUpdatePassportModal(){
    this.displayUpdatePassportModel="none"
 }
  updateView(){
    this.obs.getHobbiesForUser(this.userID).subscribe({
      next:(data:any)=>{
      this.hobbies=data
      if(this.hobbies.length!=0){
      this.isHobbyData=true
      }
      console.log("Hobbies data -",data)}
    })
    this.obs.getUserDetails(this.userID).subscribe({
      next:(data:any)=>{
        this.userName=data.FirstName+" "+data.LastName
        this.courses=data.Courses //user enrolled courses
        this.passport=data.Passport
        if (this.courses.length!=0)
        {
        this.isCourseData=true
        this.userCourses=this.coursesInDB
            for(let i=0;i<this.courses.length;i++){
              for(let j=0;j<this.userCourses.length;j++){
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
    this.myGroup=new FormGroup({
      updatedPassportID:new FormControl('',[Validators.required,Validators.min(1)]),
      updateExpiryDateForPassport:new FormControl('')
    })
    this.getAllDBCourses()
    this.updateView()
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
      // console.log("Need to call api ",this.newPassportID,this.newExpiryDateForPassport.day+"-"+this.newExpiryDateForPassport.year)
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

import { Component, Input, OnInit } from '@angular/core';
import { ObsService } from '../myservice/obs.service';
import { FormControl, FormGroup, Validators, ValidationErrors, AbstractControl, Form, FormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Passport } from '../models/passport';
import { Course } from '../models/course';
import { Hobby } from '../models/hobby';
import { User } from '../models/user';
@Component({
  selector: 'app-user-detail',
  templateUrl: './user-detail.component.html',
  styleUrls: ['./user-detail.component.css']
})
export class UserDetailComponent implements OnInit {
  addPassportForm!:FormGroup
  userName!: string
  myGroup!: FormGroup
  date: any
  newExpiryDateForPassport!: any
  userID: any
  hobbyName: string = ""
  coursesInDB: Course[] = []
  userCourses: Course[] = []
  hobbies: Hobby[] = []
  courses: Course[] = []
  passport!: Passport
  SelectedCourse: string = ""
  isCourseData: boolean = false
  isPassportData: boolean = false
  isHobbyData: boolean = false
  display = "none";
  displayPassport = "none";
  displayUpdatePassportModel = "none";
  updateExpiryDateForPassport: any
  constructor(private obs: ObsService, private router: Router,
    private activatedRoute: ActivatedRoute) {
    this.activatedRoute.paramMap.subscribe(params => {
      this.userID = params.get('userId') 
    })
  }
  goToProfile() {
    this.router.navigate(["profile/", this.userID])
  }
  goToCourse() {
    this.router.navigate(["courses"])
  }
  openUpdatePassportModel(passport: Passport) {
    console.log("Modal opened to update passport ", passport.ID)
    this.displayUpdatePassportModel = "block"
  }
  updatePassport(myGroup: FormGroup) {
    let updatedExpiryDate = ""
    this.displayUpdatePassportModel = "none"
    if (myGroup.value.updateExpiryDateForPassport != "") {
      updatedExpiryDate = myGroup.value.updateExpiryDateForPassport.year + "-" + myGroup.value.updateExpiryDateForPassport.month + "-" + myGroup.value.updateExpiryDateForPassport.day
    }
    let updatedPassport= new Passport()
    updatedPassport.ID=this.passport.ID
    updatedPassport.PassportID=myGroup.value.updatedPassportID
    updatedPassport.ExpiryDate= updatedExpiryDate
    console.log("Passport ID to be updated ", myGroup.value.updatedPassportID, updatedExpiryDate, this.passport.ID)
    this.obs.updatePassport(updatedPassport).subscribe({
      next: (data) => {
        this.updateView()
        this.displayPassport = "none"
      },
      error: (err) => {
        this.displayPassport = "none"
        alert("A user with same passport ID already exists")
        console.log("Error updating passport ", err)
      },
    })
  }
  closeUpdatePassportModal() {
    this.displayUpdatePassportModel = "none"
  }
  updateView() {
    this.obs.getHobbiesForUser(this.userID).subscribe({
      next: (data: any) => {
        this.hobbies = data
        if (this.hobbies.length != 0) {
          this.isHobbyData = true
        }
        console.log("Hobbies data -", data)
      }
    })
    this.obs.getUserDetails(this.userID).subscribe({
      next: (data: any) => {
        console.log("Getting user details ",this.coursesInDB.length)
        this.userName = data.FirstName + " " + data.LastName
        this.courses = data.Courses //user enrolled courses
        this.passport = data.Passport
        if (this.courses.length != 0) {
          console.log("Within if of course enroll")
          this.isCourseData = true
          this.userCourses = this.coursesInDB
          for (let i = 0; i < this.courses.length; i++) {
            for (let j = 0; j < this.userCourses.length; j++) {
              if (this.userCourses[j].Name === this.courses[i].Name) {
                console.log("same course in to be enrolled courses found ")
                this.userCourses.splice(j, 1)
              }
            }
            console.log("User courses ", this.userCourses)
          }
        }
        else {
          this.isCourseData=false
          console.log("Within else of course enroll",this.coursesInDB.length)
          this.userCourses=this.coursesInDB
        }
        console.log("User ca enroll in ",this.userCourses)
        if (this.passport.PassportID != 0) {
          this.isPassportData = true
        } else {
          this.isPassportData = false
        }
        console.log("Is passport data ", this.isPassportData)
        console.log("Is courses data ", this.isCourseData)
        console.log("Course data -", data.Courses)
        console.log("Passport data", data.Passport)
      },
      error: (err) => {
        console.log("Error in getting course or passport data ", err)
      }
    })
  }
  ngOnInit(): void {
    this.myGroup = new FormGroup({
      updatedPassportID: new FormControl('', [Validators.required, Validators.min(100)]),
      updateExpiryDateForPassport: new FormControl('')
    })
    this.addPassportForm= new FormGroup({
      addPassportID: new FormControl('', [Validators.required, Validators.min(100)]),
      addExpiryDateForPassport: new FormControl('',[Validators.required])
    })
    this.getAllDBCourses()
    this.updateView()
  }
  getAllDBCourses() {
    this.obs.getAllCourses().subscribe({
      next: (data: any) => {
        console.log(" DATA COURSE FROM API ", data)
        for (let i = 0; i < data.length; i++) {
          this.coursesInDB.push(data[i])
        }
      }
    }
    )
  }
  deleteHobby(hobby: Hobby) {
    console.log("HobbyID ", hobby)
    let hobbyToBeDeleted = new Hobby()
    hobbyToBeDeleted.ID=hobby.ID
    this.obs.deleteHobbyForUser(hobbyToBeDeleted).subscribe((data) => {
      console.log("Data ", data)
      this.updateView()
    })
  }
  deleteEnrolledCourse(course: Course) {
    console.log("Delete course for user ", course)
    let courseToBeUnEnrolled= new Course()
    courseToBeUnEnrolled.ID=course.ID
    this.obs.deleteCourseForUser(this.userID,courseToBeUnEnrolled).subscribe({
      next: (data) => {
        console.log("Calling update view in delete course ")
        this.updateView()
        this.userCourses.push(course)
      },
      error: (err) => {
        console.log("Error unenrolling course for user ", err)
      }
    })
  }
  openModal() {
    this.display = "block";
  }
  openModalForPassport() {
    this.displayPassport = "block";
  }
  onCloseHandledForPassport() {
    this.displayPassport = "none";
  }
  addPassport(form:FormGroup){
    if (this.isPassportData == false) {
      // console.log("Need to call api ",this.newPassportID,this.newExpiryDateForPassport.day+"-"+this.newExpiryDateForPassport.year)
      this.newExpiryDateForPassport = form.value.addExpiryDateForPassport.year + "-" + form.value.addExpiryDateForPassport.month + "-" + form.value.addExpiryDateForPassport.day
      console.log("Need to call api  -- ",form.value.addPassportID , this.newExpiryDateForPassport)
      let newPassport = new Passport()
      newPassport.PassportID=form.value.addPassportID
      newPassport.ExpiryDate=this.newExpiryDateForPassport
      this.obs.addPassportDetailsForUser(this.userID,newPassport).subscribe({
        next: (data) => {
          console.log("Passport details added ",data)
          this.updateView()
        },
        error: (err) => {
          alert(err.error)
          console.log("Error in add passport detail ", err)
        }
      })
    }
  }
  onCloseHandled() {
    console.log("New hobby name ", this.hobbyName)
    let newHobby = new Hobby()
    newHobby.HobbyName=this.hobbyName
    if (this.hobbyName != "") {
      this.obs.addUserHobby(this.userID, newHobby).subscribe({
        next: (data) => {
          this.updateView()
        },
        error: (err) => {
          console.log("Error in new add hobby ", err)
        }
      })
    }
    this.display = "none";
  }
  onCourseChange(event: any) {
    this.SelectedCourse = event.target.value
    console.log("Course add --- ", this.SelectedCourse)
    if (this.SelectedCourse !== "default") {
      for (let c of this.coursesInDB) {
        if (c.Name == this.SelectedCourse) {

          console.log("ID of selected course to be added ", c.ID)
          let courseToBeEnrolled =new Course()
          courseToBeEnrolled.ID=c.ID
          this.obs.enrollUserCourse(this.userID, courseToBeEnrolled).subscribe({
            next: (data) => {
              this.updateView()
            },
            error: (err) => {
              console.log("Error adding course ", err)
            }
          })
          alert("Course added !!")
        }
      }
    }
  }
}

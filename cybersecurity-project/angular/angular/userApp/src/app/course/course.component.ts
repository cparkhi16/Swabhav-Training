import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ApiService } from '../api/api.service';
import { Course } from '../models/course';
import { User } from '../models/user';

@Component({
  selector: 'app-course',
  templateUrl: './course.component.html',
  styleUrls: ['./course.component.css']
})
export class CourseComponent implements OnInit {
  courses!:Course[];
  message:any={text:"",status:""};
  editData:any={mode:"",course:Course};
  courseEditForm!:FormGroup;
  availableCourses!:Course[];
  currentUser!:User;
  dropdownCourses!:Course[];

  constructor(private api:ApiService, private router:Router) {
    let token=this.api.decryptData(localStorage.getItem('token')!);
    if(token!=undefined){
      console.log(token);
      this.api.checkToken(token).subscribe((data)=>{
        console.log(data);
      },(error)=>{
        this.router.navigate(["/login"]);
      });
    }else{
      this.router.navigate(["/login"]);
    }

    this.getCourses();
    //console.log(this.courses);

    this.api.getAllCourses().subscribe((data)=>{
      this.availableCourses=data;
    },(error)=>{
      this.message={text:"Error in getting courses",status:"error"};
    });

    this.courseEditForm=new FormGroup({
      'name':new FormControl('',[Validators.required,Validators.maxLength(20)])
    });
  }

  cancelEdit(){
    this.editData.mode="";
  }

  showAddCourseForm(){
    this.editData.mode="Add";
    this.dropdownCourses=[];
    //console.log(this.courses,this.availableCourses,this.dropdownCourses);
    this.dropdownCourses= this.availableCourses.filter(({ ID: id1 }) => !this.courses.some(({ ID: id2 }) => id2 === id1));
  }

  submitForm(){
    this.addCourse();
  }

  addCourse(){
    let updatedUser=new User();
    updatedUser.ID=this.currentUser.ID;
    let newCourse=new Course();
    newCourse.Name=this.courseEditForm.get('name')?.value;
    updatedUser.Courses=[newCourse];
    //console.log(this.courseEditForm.get('name')?.value);
    this.api.updateUser(updatedUser).subscribe((data)=>{
      this.message={text:"successfully added course",status:"success"};
      this.getCourses();
    },(error)=>{
      this.message={text:"Error in adding course",status:"error"};
    })
    this.editData.mode="";
  }

  deleteCourse(courseId:string){
    this.api.deleteUserCourse(this.api.decryptData(localStorage.getItem('userId')!),courseId).subscribe((data)=>{
      this.message={text:"successfully deleted course",status:"success"};
      this.getCourses();
    },
    (error)=>{
      this.message={text:"Error in deleting course",status:"error"};
    });
  }

  cancel(){
    this.message.status="";
  }

  getCourses(){
    let userId=this.api.decryptData(localStorage.getItem('userId')!);
    this.api.getUser(userId!).subscribe((data)=>{
      this.currentUser=data;
      this.courses=data.Courses!;
      //console.log(data);
    },(error)=>{
      this.message={text:"Error in getting courses",status:"error"};
    })
  }

  ngOnInit(): void {
  }

}

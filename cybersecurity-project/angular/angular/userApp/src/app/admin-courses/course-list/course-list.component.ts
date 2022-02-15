import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ApiService } from 'src/app/api/api.service';
import { Course } from 'src/app/models/course';
import { User } from 'src/app/models/user';

@Component({
  selector: 'app-course-list',
  templateUrl: './course-list.component.html',
  styleUrls: ['./course-list.component.css']
})
export class CourseListComponent implements OnInit {
  courses!:Course[];
  message:any={text:"",status:""};
  editData:any={mode:"",course:Course};
  courseEditForm!:FormGroup;

  constructor(private api:ApiService,private router:Router) {
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
      this.courses=data;
    },(error)=>{
      this.message={text:"Error in getting courses",status:"error"};
    });

    this.courseEditForm=new FormGroup({
      'name':new FormControl('',[Validators.required,Validators.maxLength(20)]),
      'price':new FormControl('',[Validators.required])
    });
  }

  editCourse(course:Course){
    this.editData.course=course;
    this.courseEditForm.get('name')?.setValue(course.Name);
    this.courseEditForm.get('price')?.setValue(course.Price);
    this.editData.mode="Update";
  }

  cancelEdit(){
    this.editData.mode="";
  }

  showAddCourseForm(){
    this.courseEditForm.setValue({name:'',price:''});
    this.editData.mode="Create";
  }

  submitForm(){
    if(this.editData.mode=="Create"){
      this.addCourse();
    }else{
      this.updateCourse();
    }
    this.editData.mode="";
  }

  updateCourse(){
    let newCourse=new Course();
    newCourse.Name=this.courseEditForm.get('name')?.value;
    newCourse.Price=this.courseEditForm.get('price')?.value;
    this.api.updateCourseById(this.editData.course.ID,newCourse).subscribe((data)=>{
      this.message={text:"successfully updated course",status:"success"};
      this.getCourses();
    },(error)=>{
      //console.log(error);
      this.message={text:"Error in updating course-"+error.error,status:"error"};
    });
    this.editData.mode="";
    this.courseEditForm.setValue({name:'',price:''});
  }

  addCourse(){
    let newCourse=new Course();
    newCourse.Name=this.courseEditForm.get('name')?.value;
    newCourse.Price=this.courseEditForm.get('price')?.value;
    this.api.createCourse(newCourse).subscribe((data)=>{
      this.message={text:"successfully added course",status:"success"};
      this.getCourses();
    },(error)=>{
      this.message={text:"Error in adding course",status:"error"};
    });
    this.editData.mode="";
    this.courseEditForm.setValue({name:'',price:''});
  }

  deleteCourse(courseId:string){
    this.api.deleteCourseById(courseId).subscribe((data)=>{
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
    this.api.getAllCourses().subscribe((data)=>{
      this.courses=data;
      //console.log(data);
    },(error)=>{
      this.message={text:"Error in getting courses",status:"error"};
    })
  }


  ngOnInit(): void {
  }

}

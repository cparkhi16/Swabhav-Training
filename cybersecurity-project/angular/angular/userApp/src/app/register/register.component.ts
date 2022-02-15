import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ApiService } from '../api/api.service';
import { Passport } from '../models/passport';
import { User } from '../models/user';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  registerForm!:FormGroup;
  model:any;
  message:any={text:"",status:""};
  constructor(private api:ApiService, private router:Router) {
    this.registerForm=new FormGroup({
      'firstName':new FormControl('',[Validators.required,Validators.maxLength(20)]),
      'lastName':new FormControl('',[Validators.required,Validators.maxLength(20)]),
      'email':new FormControl('',[Validators.required,Validators.email,Validators.maxLength(20)]),
      'password':new FormControl('',[Validators.required,Validators.maxLength(20)]),
      'country':new FormControl('',[Validators.required,Validators.maxLength(20)]),
      'expiryDate':new FormControl('',[Validators.required]),
      'levelbiba':new FormControl('',[Validators.required,Validators.max(3),Validators.min(1)]),
      'levelbell':new FormControl('',[Validators.required,Validators.max(3),Validators.min(1)]),
      'secretAnswer':new FormControl('',[Validators.required,Validators.maxLength(20)]),
    });
    //this.registerForm.get('expiryDate')?.setValue({day:2,month:1,year:2021});  
  }

  register(){
    let newUser=new User();
    newUser.FirstName=this.registerForm.get('firstName')?.value;
    newUser.LastName=this.registerForm.get('lastName')?.value;
    newUser.Email=this.registerForm.get('email')?.value;
    newUser.Password=this.registerForm.get('password')?.value;
    newUser.LevelBIBA=this.registerForm.get('levelbiba')?.value;
    newUser.LevelBell=this.registerForm.get('levelbell')?.value;
    newUser.SecretAnswer=this.registerForm.get('secretAnswer')?.value;
    newUser.Passport=new Passport();
    newUser.Passport.Country=this.registerForm.get('country')?.value;
    let expiryDate=this.registerForm.get('expiryDate')?.value
    console.log(expiryDate);
    newUser.Passport.ExpiryDate=expiryDate.year+"-"+expiryDate.month+"-"+expiryDate.day;
    console.log(newUser);
    this.api.createUser(newUser).subscribe((data)=>{
      this.message={text:"Successfully added user",status:"success"};
      this.router.navigate(["/login"]);
    },(error)=>{
      this.message={text:"Error in adding user",status:"error"};
    });
  }

  cancel(){
    this.message.status="";
  }

  ngOnInit(): void {
  }

}

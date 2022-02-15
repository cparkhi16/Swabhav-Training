import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, ValidationErrors, ValidatorFn, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ApiService } from '../api/api.service';
import { LoginData } from '../models/loginData';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginForm!:FormGroup;
  wrongCredentials:boolean=false;

  constructor(private api:ApiService, private router:Router) { 
    let token=this.api.decryptData(localStorage.getItem('token')!);
    let userId=this.api.decryptData(localStorage.getItem('userId')!);
    if(token!=undefined){
      console.log(token);
      this.api.checkToken(token).subscribe((data)=>{
        console.log(data);
        this.router.navigate(["/dashboard/"+userId+"/passport"]);
      },(error)=>{
        console.log(error);
      });
    }
    this.loginForm=new FormGroup({
      'email':new FormControl('',[Validators.required,Validators.email,Validators.maxLength(20)]),
      'password':new FormControl('',[Validators.required,Validators.maxLength(10),this.createPasswordStrengthValidator()]),
      'secretAnswer':new FormControl('',[Validators.required,Validators.maxLength(20)]),
    });
  }

  ngOnInit(): void {
  }
  onLogin(){
    console.log(this.loginForm.get('password')?.errors);
    console.log(this.loginForm.get('email')?.value, this.loginForm.get('password')?.value);
    this.api.login(this.loginForm.get('email')?.value, this.loginForm.get('password')?.value,this.loginForm.get('secretAnswer')?.value).subscribe((data:LoginData)=>{
      console.log(data);
      localStorage.setItem('token', this.api.encryptData(data.token!)!);
      localStorage.setItem('userId',this.api.encryptData(data.userId!)!);
      this.router.navigate(["/dashboard/"+data.userId+"/passport"]);
    },(error)=>{
      console.log(error);
      this.wrongCredentials=true;
    })
  }

  createPasswordStrengthValidator(): ValidatorFn {
    return (control:AbstractControl) : ValidationErrors | null => {
        const value = control.value;
        if (!value) {
            return null;
        }
        //const hasUpperCase = /[A-Z]+/.test(value);
        const hasLowerCase = /[a-z]+/.test(value);
        //const hasNumeric = /[0-9]+/.test(value);
        const passwordValid = hasLowerCase;
        return !passwordValid ? {passwordStrength:true}: null;
    }
  }

}

import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, ValidationErrors, ValidatorFn, Validators } from '@angular/forms';
import { ServiceService } from '../service/service.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginForm!:FormGroup;
  isUserLoggedIn:boolean=false;

  constructor(private myService:ServiceService) { 
    let token=localStorage.getItem('token');
    if(token!=undefined){
      console.log(token);
      this.myService.checkToken(token).subscribe((data)=>{
        console.log(data);
        this.isUserLoggedIn=true;
      },(error)=>{
        console.log(error);
        this.isUserLoggedIn=false;
      });
    }
    this.loginForm=new FormGroup({
      'email':new FormControl('',[Validators.required,Validators.email,Validators.maxLength(20)]),
      'password':new FormControl('',[Validators.required,Validators.maxLength(10),this.createPasswordStrengthValidator()]),
      'passportNo':new FormControl('',[Validators.required,this.passportNoValidator()])
    });
  }

  ngOnInit(): void {
  }

  onLogin(){
    // let e=this.loginForm.get('password')?.errors;
    //console.log(this.loginForm.get('password')?.errors?['maxLength']['requiredLength']);
    console.log(this.loginForm.get('password')?.errors);
    console.log(this.loginForm.get('email')?.value, this.loginForm.get('password')?.value);
    this.myService.login(this.loginForm.get('email')?.value, this.loginForm.get('password')?.value).subscribe((data)=>{
      console.log(data);
      localStorage.setItem('token', data.token);
      this.isUserLoggedIn=true;
    },(error)=>{
      console.log(error);
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

passportNoValidator():ValidatorFn{
  return (control:AbstractControl):ValidationErrors | null =>{
    const value = control.value;
      if (!value) {
          return null;
      }
      const hasCorrectFormat = /^[A-Z]{3}[0-9]{7}$/.test(value);
      const passwordValid = hasCorrectFormat;
      return !passwordValid ? {passportInvalid:true}: null;
  }
}

}

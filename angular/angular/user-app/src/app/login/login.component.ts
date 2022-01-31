import { Component, DoCheck, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators,ValidationErrors,AbstractControl } from '@angular/forms';
import { ObsService } from '../myservice/obs.service';
import { ValidatorFn } from '@angular/forms';
import { ActivatedRoute, Router} from '@angular/router';
import { User } from '../models/user';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  userid:any
  myForm: FormGroup;
  isValidToken:any
  constructor(private obs:ObsService,private router:Router,private route: ActivatedRoute) { 
    this.myForm = new FormGroup({
      email: new FormControl('',[Validators.email,Validators.required,Validators.maxLength(50)]),
      password: new FormControl('',[Validators.required,Validators.maxLength(10),Validators.min(5)]),
      passport: new FormControl('',[this.passportValidator()])
    });
  }

  ngOnInit(): void {
  }
  onSubmit(form:FormGroup){
    console.log("Is form valid ",form.valid)
    let myData={
      Email:form.value.email,
      Password:form.value.password
    }
   console.log("Mydata ",JSON.stringify(myData))
   let user = new User()
   user.Email=form.value.email
   user.Password=form.value.password
    this.obs.validateUser(user).subscribe({
      next:(data)=>{console.log("Data from login ",data)
      this.userid=data.ID
      console.log("User ID in parent ",this.userid)
      let token:any=this.obs.encryptData(data.Token)
      console.log("Encrypted token ",token)
      let decryptedToken =this.obs.decryptData(token)
      console.log("Decrypted token ",decryptedToken)
      localStorage.setItem('Token', token);
      this.isValidToken=true
      this.router.navigate(['userDetail/',this.userid]);
    },
      error:(err)=>{console.log("Error ",err)
    alert(err.error)}
    })  
  }

  
passportValidator(): ValidatorFn {
  return (control:AbstractControl) : ValidationErrors | null => {
      const value = control.value;
      if (!value) {
          return null;
      }
      if (value.length!=10){
        return {inValidPassport:true}
      }
      //let check= /^[a-zA-Z0-9]+$/.test(value)
      let first=value.substring(0,3)
      let last=value.substring(3,value.length)
      const checkLast=/^[0-9]+$/.test(last)
      const checkFirstThree=/^[a-zA-Z]+$/.test(first)
      console.log("Checking ",checkFirstThree)
      console.log("First 3 chars ",first)
      console.log("Check last 7 ",checkLast)
      const passportValid = checkFirstThree && checkLast;
      console.log("is passport valid ",passportValid)
      return !passportValid ?{inValidPassport:true}:null
      
  }
}
}

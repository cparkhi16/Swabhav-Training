import { Component, DoCheck, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators,ValidationErrors,AbstractControl } from '@angular/forms';
import { ObsService } from '../myservice/obs.service';
import { ValidatorFn } from '@angular/forms';
import { ActivatedRoute, Router} from '@angular/router'
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit,DoCheck {
  userid:any
  myForm: FormGroup;
  isValidToken:any
  constructor(private obs:ObsService,private router:Router,private route: ActivatedRoute) { 
    //localStorage.setItem('Token',"ABC")
    this.myForm = new FormGroup({
      email: new FormControl('',[Validators.email,Validators.required,Validators.maxLength(50)]),
      password: new FormControl('',[Validators.required,Validators.maxLength(10),Validators.min(5)]),
      passport: new FormControl('',[this.passportValidator()])
    });
  }

  ngOnInit(): void {
    let currentToken =localStorage.getItem('Token')
    this.obs.validateToken(currentToken).subscribe({
      next:(data)=>
      {
        //let d =JSON.parse(data)
        console.log("Valid token data ",data)
        //this.isValidToken=data.IsValidToken
        //this.isValidToken=d.IsValidToken}
      },

      error:(err)=>{
        console.log("Error validating token ",err)
      }
    })
  }
  ngDoCheck(): void {
    //   let currentToken =localStorage.getItem('Token')
    // this.obs.validateToken(currentToken).subscribe({
    //   next:(data)=>
    //   {
    //     //let d =JSON.parse(data)
    //     console.log("Valid token data ",data)
    //     this.isValidToken=data.IsValidToken
    //     //this.isValidToken=d.IsValidToken}
    //   },
    //   error:(err)=>{
    //     console.log("Error validating token ",err)
    //   }
    // })
  }
  regsiter():void{
    this.isValidToken=true
  }
  onSubmit(form:FormGroup){
    console.log("Is form valid ",form.valid)
    let myData={
      Email:form.value.email,
      Password:form.value.password
    }
   console.log("Mydata ",JSON.stringify(myData))
    this.obs.validateUser(form.value.email,form.value.password).subscribe({
      next:(data)=>{console.log("Data from login ",data)
      this.userid=data.ID
      console.log("User ID in parent ",this.userid)
      localStorage.setItem('Token', data.Token);
      this.isValidToken=true
      this.router.navigate(['userDetail/',this.userid], {relativeTo:this.route});
    },
      error:(err)=>{console.log("Error ",err)}
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

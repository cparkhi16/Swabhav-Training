import { Component, OnInit } from '@angular/core';
import { FormGroup,Validators,FormControl } from '@angular/forms';
import { ObsService } from '../myservice/obs.service';
import { ActivatedRoute, Router} from '@angular/router'
import { User } from '../models/user';
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  isValidToken!:boolean
  userid!:string
  myForm!:FormGroup
  constructor(private obs:ObsService,private router:Router,private route: ActivatedRoute) {
    this.isValidToken=false
   }

  ngOnInit(): void {
    this.myForm = new FormGroup({
      email: new FormControl('',[Validators.email,Validators.required,Validators.maxLength(50)]),
      password: new FormControl('',[Validators.required,Validators.maxLength(10),Validators.min(5)]),
      firstname:new FormControl('',[Validators.required,Validators.maxLength(30)]),
      lastname:new FormControl('',[Validators.required,Validators.maxLength(30)]),
      address:new FormControl('',[Validators.required,Validators.maxLength(60)])
    });
  }
  onSubmit(form:FormGroup){
    let newUser= new User()
    newUser.Email=form.value.email
    newUser.FirstName=form.value.firstname
    newUser.LastName=form.value.lastname
    newUser.Password=form.value.password
    newUser.Address=form.value.address
    this.obs.createUser(newUser).subscribe({
      next:(data)=>{console.log("Data from login ",data)
      this.userid=data.UserID
      console.log("User ID in register ",this.userid)
      let token:any=this.obs.encryptData(data.Token)
      localStorage.setItem('Token', token);
      this.isValidToken=true
      this.router.navigate(['userDetail/',this.userid]);
    },
      error:(err)=>{
        alert(err.error)
        console.log("Error ",err)}
    })
  }
}

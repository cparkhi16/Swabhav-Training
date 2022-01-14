import { Component, OnInit } from '@angular/core';
import { FormGroup,Validators,FormControl } from '@angular/forms';
import { ObsService } from '../myservice/obs.service';
import { ActivatedRoute, Router} from '@angular/router'
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  isValidToken:any
  userid:any
  myForm:any
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
  onSubmit(form:any){
    this.obs.createUser(form.value.email,form.value.password,form.value.firstname,form.value.lastname,form.value.address).subscribe({
      next:(data)=>{console.log("Data from login ",data)
      this.userid=data.UserID
      console.log("User ID in register ",this.userid)
      localStorage.setItem('Token', data.Token);
      this.isValidToken=true
      this.router.navigate(['userDetail/',this.userid], {relativeTo:this.route});
    },
      error:(err)=>{console.log("Error ",err)}
    })
  }
}

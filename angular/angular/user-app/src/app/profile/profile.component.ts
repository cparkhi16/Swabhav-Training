import { ObsService } from './../myservice/obs.service';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormControl, FormGroup, Validators,ValidationErrors,AbstractControl } from '@angular/forms';
@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  updateForm!:FormGroup
  user:any
  userID:any
  displayUpdateProfileModal="none"
  constructor(private obs:ObsService,private router:Router,
    private activatedRoute:ActivatedRoute) { 
      this.activatedRoute.paramMap.subscribe(params=>{
        this.userID=params.get('userId') //+ string to number
      })
    }
    getUserDetails(){
      this.obs.getUserDetails(this.userID).subscribe({
        next:(data:any)=>{
          this.user=data
          console.log("User data ",this.user)
        }
      })
    }
  ngOnInit(): void {
    console.log(this.userID)
    this.getUserDetails()
    //let userFirstName:string=this.user.FirstName
     this.updateForm=new FormGroup({
            updatedFirstName:new FormControl('',[Validators.maxLength(30)]),
            updatedLastName:new FormControl('',[Validators.maxLength(30)]),
            updatedAddress:new FormControl('',[Validators.maxLength(50)]),
            updatedPassword:new FormControl('',[Validators.maxLength(30)]),
            updatedEmail:new FormControl('',[Validators.email])
          })
  }
  openUpdateProfileModal(){
    this.displayUpdateProfileModal="block"
  }
  closeUpdateModal(){
    this.displayUpdateProfileModal="none"
  }
  updateUser(form:any){
    console.log("Updated details ",form.value)
    this.obs.updateUserProfile(this.userID,form.value.updatedAddress,form.value.updatedPassword,form.value.updatedFirstName,form.value.updatedLastName,form.value.updatedEmail).subscribe({
      next:(data)=>{
        this.getUserDetails()
      },error:(err)=>{
        console.log(err)
      }
    })
    this.displayUpdateProfileModal="none"
  }
  logout(){
    localStorage.setItem('Token',"");
    this.router.navigate(['login']);
  }
  goToUserDetail(){
    this.router.navigate(["userDetail/",this.userID])
  }
}

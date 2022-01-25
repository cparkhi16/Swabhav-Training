import { ObsService } from './../myservice/obs.service';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormControl, FormGroup, Validators,ValidationErrors,AbstractControl } from '@angular/forms';
import { User } from '../models/user';
@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  updateForm!:FormGroup
  user!:User
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
      //this.updateForm.setValue
  }
  openUpdateProfileModal(){
    this.displayUpdateProfileModal="block"
    this.updateForm.get('updatedFirstName')?.setValue(this.user.FirstName)
    this.updateForm.get('updatedLastName')?.setValue(this.user.LastName)
    this.updateForm.get('updatedAddress')?.setValue(this.user.Address)
    this.updateForm.get('updatedEmail')?.setValue(this.user.Email)
  }
  closeUpdateModal(){
    this.displayUpdateProfileModal="none"
  }
  updateUser(form:any){
    console.log("Updated details ",form.value)
    let updatedUserDetails= new User()
    updatedUserDetails.ID=this.userID
    updatedUserDetails.Address=form.value.updatedAddress
    updatedUserDetails.Password=form.value.updatedPassword
    updatedUserDetails.FirstName=form.value.updatedFirstName
    updatedUserDetails.LastName=form.value.updatedLastName
    updatedUserDetails.Email=form.value.updatedEmail
    this.obs.updateUserProfile(updatedUserDetails).subscribe({
      next:(data)=>{
        this.getUserDetails()
      },error:(err)=>{
        alert(err.error)
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
  deleteAccount(){
    this.obs.deleteUser(this.user).subscribe({
      next:(data)=>{
        this.logout()
      },
      error:(err)=>{
        alert("Error deleting user")
        console.log("Error deleting user ",err)      
      }
    })
  }
}

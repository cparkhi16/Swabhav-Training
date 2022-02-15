import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ApiService } from '../api/api.service';
import { Passport } from '../models/passport';

@Component({
  selector: 'app-passport',
  templateUrl: './passport.component.html',
  styleUrls: ['./passport.component.css']
})
export class PassportComponent implements OnInit {
  passports:Passport[]=[];
  message:any={text:"",status:""};
  editData:any={mode:"",passport:Passport};
  passportEditForm:FormGroup;
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
    this.getPassports();
    this.passportEditForm=new FormGroup({
      'passportNo':new FormControl('',[Validators.required,Validators.maxLength(36)]),
      'expiryDate':new FormControl('',[Validators.required]),
      'country':new FormControl('',[Validators.required,Validators.maxLength(20)])
    });
  }

  editPassport(passport:Passport){
    console.log(passport);
    this.editData.passport=passport;
    this.passportEditForm.get('passportNo')?.setValue(passport.PassportNo);
    let expiryDate=passport.ExpiryDate;
    this.passportEditForm.get('expiryDate')?.setValue({year:parseInt(expiryDate?.split("-")[0]!),month:parseInt(expiryDate?.split("-")[1]!),day:parseInt(expiryDate?.split("-")[2].substring(0,2)!)});
    //console.log({year:parseInt(expiryDate?.split("-")[0]!),month:parseInt(expiryDate?.split("-")[1]!),day:parseInt(expiryDate?.split("-")[2].substring(0,2)!)});
    this.passportEditForm.get('country')?.setValue(passport.Country);
    this.editData.mode="Update";
  }

  updatePassport(){
    let expiryDate=this.passportEditForm.get('expiryDate')?.value
    let formattedExpiryDate=expiryDate.year+"-"+expiryDate.month+"-"+expiryDate.day;
    let passport:Passport={PassportNo:this.passportEditForm.get('passportNo')?.value, ExpiryDate:formattedExpiryDate, Country:this.passportEditForm.get('country')?.value};
    this.api.updatePassportById(this.editData.passport.ID,passport).subscribe((data)=>{
      console.log("updated passport");
      this.message={text:"Successfully updated passport",status:"success"};
      this.getPassports();
    },(error)=>{
      this.message={text:"Error in updating passport",status:"error"};
    });
    this.editData.mode="";
    this.editData.passportId="";
  }

  cancelEdit(){
    this.editData.mode="";
    this.editData.passportId="";
  }


  deletePassport(passportId:string){
    this.api.deletePassportById(passportId).subscribe((data)=>{
      console.log("deleted passport");
      this.message={text:"Successfully deleted passport",status:"success"};
      this.getPassports();
    },(error)=>{
      this.message={text:"Error in deleting passport",status:"error"};
    })
  }

  cancel(){
    this.message.status="";
  }

  getPassports(){
    let userId=this.api.decryptData(localStorage.getItem('userId')!);
    this.api.getPassportByUserId(userId!).subscribe((data:any)=>{
      this.passports=data;
      console.log(data);
    });
  }


  ngOnInit(): void {
  }

}

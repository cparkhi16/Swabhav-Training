import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { map, Observable } from 'rxjs';
import { ApiService } from '../api/api.service';
import { Hobby } from '../models/hobby';

@Component({
  selector: 'app-hobby',
  templateUrl: './hobby.component.html',
  styleUrls: ['./hobby.component.css']
})
export class HobbyComponent implements OnInit {
  hobbies:Hobby[]=[];
  message:any={text:"",status:""};
  editData:any={mode:"",hobby:Hobby};
  hobbyEditForm:FormGroup;
  constructor(private api:ApiService, private router:Router) { 
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
    this.getHobbies();
    this.hobbyEditForm=new FormGroup({
      'name':new FormControl('',[Validators.required,Validators.maxLength(20)])
    });
  }

  editHobby(hobby:Hobby){
    console.log(hobby);
    this.editData.hobby=hobby;
    this.hobbyEditForm.get('name')?.setValue(hobby.Name);
    this.editData.mode="Update";
  }

  updateHobby(){
    let hobby:Hobby={Name:this.hobbyEditForm.get('name')?.value}
    this.api.updateByHobbyId(this.editData.hobby.ID,hobby).subscribe((data)=>{
      console.log("updated hobby");
      this.message={text:"Successfully updated hobby",status:"success"};
      this.getHobbies();
    },(error)=>{
      this.message={text:"Error in updating hobby",status:"error"};
    });
    this.editData.mode="";
    this.editData.hobbyId="";
    this.hobbyEditForm.setValue({name:''});
  }

  cancelEdit(){
    this.editData.mode="";
    this.editData.hobbyId="";
  }

  showAddHobbyForm(){
    this.editData.mode="Create";
    this.editData.hobbyId="";
    this.hobbyEditForm.setValue({name:''});
  }

  submitForm(){
    if(this.editData.mode=="Update"){
      this.updateHobby();
    }else{
      this.addHobby();
    }
  }

  addHobby(){
    let hobby:Hobby={Name:this.hobbyEditForm.get('name')?.value,UserId:this.api.decryptData(localStorage.getItem('userId')!)};
    this.api.addHobby(hobby).subscribe((data)=>{
      this.message={text:"Successfully added hobby",status:"success"};
      this.getHobbies();
    },(error)=>{
      this.message={text:"Error in adding hobby",status:"error"};
    });
    this.editData.mode="";
    this.editData.hobbyId="";
    this.hobbyEditForm.setValue({name:''});
  }

  deleteHobby(hobbyId:string){
    this.api.deleteHobbyById(hobbyId).subscribe((data)=>{
      console.log("deleted hobby");
      this.message={text:"Successfully deleted hobby",status:"success"};
      this.getHobbies();
    },(error)=>{
      this.message={text:"Error in deleting hobby",status:"error"};
    })
  }

  cancel(){
    this.message.status="";
    this.hobbyEditForm.setValue({name:''});
  }

  getHobbies(){
    let userId=this.api.decryptData(localStorage.getItem('userId')!);
    this.api.getHobbiesByUserId(userId!).subscribe((data:any)=>{
      this.hobbies=data;
      console.log(data);
    });
  }

  ngOnInit(): void {
  }

}

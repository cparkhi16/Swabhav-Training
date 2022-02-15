import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiService } from '../api/api.service';
import { File } from '../models/file';

@Component({
  selector: 'app-files',
  templateUrl: './files.component.html',
  styleUrls: ['./files.component.css']
})
export class FilesComponent implements OnInit {
  files:any={"ReadList":[],"WriteList":[]};
  //private completeFileList:any=[];
  userId:string;
  displayEditModal:string="none";
  displayWriteModal:string="none";
  fileData:string="yyyy";
  writeForm!:FormGroup;
  writeFileInfo!:File;
  message:any={text:"",status:""};
  constructor(private api:ApiService) {
    this.userId=this.api.decryptData(localStorage.getItem('userId')!);
    this.api.getAccessibleFilesOfUser(this.userId).subscribe(data=>{
      console.log(data);
      this.files=data;
      console.log(this.files.ReadList);
    },error=>{
      console.log(error);
    });
    this.writeForm=new FormGroup(
      {
        'data':new FormControl('',[Validators.required])
      }
    )
  }

  readFile(fileId:string):string{
    this.api.readFile(this.userId,fileId).subscribe((data)=>{
      console.log(data);
      this.displayEditModal = "block";
      return data;
    },(error)=>{
      console.log(error);
      return "";
    });
    return "";
  }

  writeFile(fileId:string,data:string){
    this.api.writeFile(this.userId,fileId,data).subscribe((data)=>{
      console.log(data);
      this.message={text:"Successfully updated file data",status:"success"};
    },(error)=>{
      console.log(error);
      this.displayWriteModal = "none";
      this.message={text:"Error in updating file data",status:"error"};
    });
  }

  openEditModal(file:File,operation:string) {
    console.log("Modal opened to update passport ", file.ID);
    if(operation==="r"){
      // this.fileData=this.readFile(file.ID!);
      this.api.readFile(this.userId,file.ID!).subscribe((data)=>{
        console.log(data);
        this.fileData=data;
        console.log(this.fileData);
        this.displayEditModal = "block";
      },(error)=>{
        console.log(error);
        this.message={text:"Error in updating file data",status:"error"};
      });
    }
    else{
      this.displayWriteModal = "block";
      this.writeFileInfo=file;
    }
  }

  writeFormSubmit(writeForm:FormGroup){
    this.writeFile(this.writeFileInfo.ID!,writeForm.value.data);
    this.writeForm.value.data="";
    this.displayWriteModal = "none";
  }

  closeEditModal(){
    this.displayEditModal = "none";
    this.fileData="";
  }

  cancel(){
    this.message.status="";
  }

  closeWriteModal(){
    this.displayWriteModal="none";
    this.writeForm.setValue({data:""});
    this.writeForm.reset();
  }

  ngOnInit(): void {
  }

}

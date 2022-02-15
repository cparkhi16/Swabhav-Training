import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { LoginData } from '../models/loginData';
import { Hobby } from '../models/hobby';
import { TokenError } from '@angular/compiler/src/ml_parser/lexer';
import { Passport } from '../models/passport';
import { Course } from '../models/course';
import { User } from '../models/user';
import * as CryptoJS from 'crypto-js';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  encryptSecretKey:string="yogesh";
  constructor(private http:HttpClient) { }

  getCourses():Observable<any>{
    return this.http.get<any>("http://localhost:8000/courses/");
  }

  login(email:string,password:string,secretAnswer:string):Observable<LoginData>{
    return this.http.post<LoginData>("http://localhost:8000/login",{"email":email,"password":password,"secretAnswer":secretAnswer});
  }

  checkToken(token:string):Observable<any>{
    return this.http.post<any>("http://localhost:8000/checkToken",{"token":token});
  }

  getHobbiesByUserId(userId:string):Observable<Hobby>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.get<Hobby>("http://localhost:8000/hobbies/"+userId,{headers:headers});
  }

  deleteHobbyById(id:string):Observable<any>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.delete<any>("http://localhost:8000/hobbies/"+id,{headers:headers});
  }

  updateByHobbyId(id:string,hobby:Hobby):Observable<any>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    console.log("hobby-",hobby);
    return this.http.put<any>("http://localhost:8000/hobbies/"+id,hobby,{headers:headers});
  }

  addHobby(hobby:Hobby):Observable<any>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.post<any>("http://localhost:8000/hobbies/",hobby,{headers:headers});
  }

  getPassportByUserId(userId:string):Observable<Hobby>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.get<Hobby>("http://localhost:8000/passports/"+userId,{headers:headers});
  }

  deletePassportById(id:string):Observable<any>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.delete<any>("http://localhost:8000/passports/"+id,{headers:headers});
  }

  updatePassportById(id:string,passport:Passport):Observable<any>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.put<any>("http://localhost:8000/passports/"+id,passport,{headers:headers});
  }

  // getUser(userId:string){
  //   let headers = new HttpHeaders().set('access_token',localStorage.getItem('token')!);
  //   this.http.get<any>("http://localhost:8000/users/"+userId,{headers:headers}).subscribe((data)=>{
  //     this.currentUser=data;
  //     console.log(data);
  //   },(error)=>{
  //     console.log(error);
  //   });
  // }

  getUser(userId:string):Observable<User>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.get<any>("http://localhost:8000/users/"+userId,{headers:headers});
  }

  createUser(user:User):Observable<User>{
    //let headers = new HttpHeaders().set('access_token',localStorage.getItem('token')!);
    return this.http.post<any>("http://localhost:8000/users/",user);
  }

  updateUser(user:User):Observable<User>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.put<any>("http://localhost:8000/users/"+user.ID,user,{headers:headers})
  }

  createCourse(course:Course):Observable<Course>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.post<any>("http://localhost:8000/courses/",course,{headers:headers});
  }

  deleteCourseById(id:string):Observable<any>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.delete<any>("http://localhost:8000/courses/"+id,{headers:headers});
  }

  updateCourseById(id:string,course:Course):Observable<any>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.put<any>("http://localhost:8000/courses/"+id,course,{headers:headers});
  }

  getAllCourses():Observable<any>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.get<any>("http://localhost:8000/courses/",{headers:headers});
  }

  deleteUserCourse(userId:string,courseId:string):Observable<any>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.delete<any>("http://localhost:8000/users/"+userId+"/course/"+courseId,{headers:headers});
  }

  getAccessibleFilesOfUser(userId:string):Observable<any>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.get<any>("http://localhost:8000/files/"+userId,{headers:headers});
  }

  readFile(userId:string,fileId:string):Observable<any>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.get<any>("http://localhost:8000/files/"+userId+"/read/"+fileId,{headers:headers});
  }

  writeFile(userId:string,fileId:string,data:string):Observable<any>{
    let headers = new HttpHeaders().set('access_token',this.decryptData(localStorage.getItem('token')!));
    return this.http.post<any>("http://localhost:8000/files/"+userId+"/write/"+fileId,{"Data":data},{headers:headers});
  }

  encryptData(data:string) {
    try {
      return CryptoJS.AES.encrypt(JSON.stringify(data), this.encryptSecretKey).toString();
    } catch (e) {
      console.log(e);
    }
    return null
  }

  decryptData(data:string) {
    try {
      const bytes = CryptoJS.AES.decrypt(data, this.encryptSecretKey);
      if (bytes.toString()) {
        return JSON.parse(bytes.toString(CryptoJS.enc.Utf8));
      }
      return data;
    } catch (e) {
      console.log(e);
    }
  }

}

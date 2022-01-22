import { Injectable } from '@angular/core';
import { Observable,interval } from 'rxjs';
import { map } from 'rxjs';
import { HttpClient, HttpParams, HttpHeaders } from '@angular/common/http';
import { Course } from '../models/course';
import { Passport } from '../models/passport';
import { Hobby } from '../models/hobby';
import { User } from '../models/user';
@Injectable({
  providedIn: 'root'
})
export class ObsService {
  obs!:Observable<any>
  baseURL: string = "http://localhost:9000/courses";
  //baseURL: string = "http://app:9000/courses";
  constructor(private http: HttpClient) { }
  getRandomInt(max:number) {
    return Math.floor(Math.random() * max);
  }
  getData():Observable<any>{
    this.obs=interval(5000).pipe(map(n=>{
      n=this.getRandomInt(36)
      var today = new Date()
      var time = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();
      return {"Data":n,"Time":time}
    }))
    return this.obs
  }
  getDataFromApi():Observable<any>{
   return this.http.get(this.baseURL, {responseType: 'text'})
  }

  getRandomDataFromApi():Observable<any>{
    return interval(5000).pipe(map(n=>{return this.http.get("http://localhost:9000/getRandomNumber", {responseType: 'text'})}))
    //return interval(5000).pipe(map(n=>{return this.http.get("http://app:9000/getRandomNumber", {responseType: 'text'})}))
  }
  validateUser(user:User):Observable<any>{
     return this.http.post<any>("http://localhost:9000/login",{"Email":user.Email,"Password":user.Password})
    //return this.http.post<any>("http://app:9000/login",{"Email":email,"Password":password})
  }
  validateToken(token:any){
    return this.http.post<any>("http://localhost:9000/validateToken",{"Token":token})
    //return this.http.post<any>("http://app:9000/validateToken",{"Token":token})
  }
  getHobbiesForUser(id:string){
   return this.http.get("http://localhost:9000/users/"+id+"/hobbies") 
   //return this.http.get("http://app:9000/users/"+id+"/hobbies") 
  }
  getUserDetails(id:string){
    return this.http.get("http://localhost:9000/users/"+id) 
    //return this.http.get("http://app:9000/users/"+id) 
   }
  deleteHobbyForUser(hobby:Hobby){ 
    let currentToken:any=localStorage.getItem('Token')
    let headers= new HttpHeaders().set('Token',currentToken);
    return this.http.delete("http://localhost:9000/hobbies/"+hobby.ID,{headers:headers}) 
    //return this.http.delete("http://app:9000/hobbies/"+id,{headers:headers}) 
  }
  createUser(user:User){
    console.log("User firstname ",user.FirstName)
    return this.http.post<any>("http://localhost:9000/users",{"Email":user.Email,"Password":user.Password,"FirstName":user.FirstName,"LastName":user.LastName,"Address":user.Address})
    //return this.http.post<any>("http://app:9000/users",{"Email":email,"Password":password,"FirstName":firstName,"LastName":lastName,"Address":address})
  }
  getAllCourses(){
    return this.http.get("http://localhost:9000/courses")
    //return this.http.get("http://app:9000/courses")
  }
  enrollUserCourse(id:any,course:Course){
    let currentToken:any=localStorage.getItem('Token')
     let headers= new HttpHeaders().set('Token',currentToken);
     console.log("Enroll course called ")
     return this.http.put<any>("http://localhost:9000/users/"+id,{"Courses":[{"ID":course.ID}]},{headers:headers})
    //return this.http.put<any>("http://app:9000/users/"+id,{"Courses":[{"ID":courseID}]},{headers:headers})
  }
  addUserHobby(id:any,hobby:Hobby){
    let currentToken:any=localStorage.getItem('Token')
     let headers= new HttpHeaders().set('Token',currentToken);
     return this.http.put<any>("http://localhost:9000/users/"+id,{"Hobbies":[{"HobbyName":hobby.HobbyName}]},{headers:headers})
    //return this.http.put<any>("http://app:9000/users/"+id,{"Hobbies":[{"HobbyName":hobbyName}]},{headers:headers})
  }
  addPassportDetailsForUser(id:any,passport:Passport){
     let currentToken:any=localStorage.getItem('Token')
     let headers= new HttpHeaders().set('Token',currentToken);
     console.log("Passport id ",Number(passport.PassportID))
    return this.http.post<any>("http://localhost:9000/users/"+id+"/passport",{"Passport":{"PassportID":Number(passport.PassportID),"ExpiryDate":passport.ExpiryDate}},{headers:headers})
    //return this.http.post<any>("http://app:9000/users/"+id+"/passport",{"Passport":{"PassportID":Number(passportID),"ExpiryDate":expiryDate}},{headers:headers})
  }
  deleteCourseForUser(id:any,course:Course){
     let currentToken:any=localStorage.getItem('Token')
     let headers= new HttpHeaders().set('Token',currentToken);
     return this.http.delete("http://localhost:9000/users/"+id+"/courses/"+course.ID,{headers:headers}) 
    //return this.http.delete("http://app:9000/users/"+id+"/courses/"+courseid,{headers:headers}) 
  }
  addCourse(course:Course){
    let currentToken:any=localStorage.getItem('Token')
    let headers= new HttpHeaders().set('Token',currentToken);
     return this.http.post<any>("http://localhost:9000/courses",{"Name":course.Name},{headers:headers}) 
    //return this.http.post<any>("http://app:9000/courses",{"Name":courseName},{headers:headers}) 
  }
  deleteCourse(course:Course){
  let currentToken:any=localStorage.getItem('Token')
  let headers= new HttpHeaders().set('Token',currentToken);
  let params = new HttpParams();
  params = params.append('hardDelete', "false");
   return this.http.delete("http://localhost:9000/courses/"+course.ID,{headers:headers,params: params})
   //return this.http.delete("http://app:9000/courses/"+id,{headers:headers,params: params}) 
  }
  deleteUser(user:User){
    let currentToken:any=localStorage.getItem('Token')
    let headers= new HttpHeaders().set('Token',currentToken);
    let params = new HttpParams();
    params = params.append('hardDelete', "false");
     return this.http.delete("http://localhost:9000/users/"+user.ID,{headers:headers,params: params})
     //return this.http.delete("http://app:9000/courses/"+id,{headers:headers,params: params}) 
    }
  updateCourse(course:Course){
    let currentToken:any=localStorage.getItem('Token')
    let headers= new HttpHeaders().set('Token',currentToken);
   return this.http.put<any>("http://localhost:9000/courses/"+course.ID,{"Name":course.Name},{headers:headers})
   // return this.http.put<any>("http://app:9000/courses/"+id,{"Name":courseName},{headers:headers})
  }
  updateHobby(hobby:Hobby){
    let currentToken:any=localStorage.getItem('Token')
    let headers= new HttpHeaders().set('Token',currentToken);
   return this.http.put<any>("http://localhost:9000/hobbies/"+hobby.ID,{"HobbyName":hobby.HobbyName},{headers:headers})
   // return this.http.put<any>("http://app:9000/courses/"+id,{"Name":courseName},{headers:headers})
  }
  updatePassport(passport:Passport){
    return this.http.put<any>("http://localhost:9000/passport/"+passport.ID,{"PassportID":Number(passport.PassportID),"ExpiryDate":passport.ExpiryDate})
    //return this.http.put<any>("http://app:9000/passport/"+id,{"PassportID":Number(passportID),"ExpiryDate":expiryDate})
  }
  updateUserProfile(user:User){
    let currentToken:any=localStorage.getItem('Token')
     let headers= new HttpHeaders().set('Token',currentToken);
     return this.http.put<any>("http://localhost:9000/users/"+user.ID,{"FirstName":user.FirstName,"LastName":user.LastName,"Email":user.Email,"Password":user.Password,"Address":user.Address},{headers:headers})
  }
}

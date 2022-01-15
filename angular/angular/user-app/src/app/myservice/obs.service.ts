import { Injectable } from '@angular/core';
import { Observable,interval } from 'rxjs';
import { map } from 'rxjs';
import { HttpClient, HttpParams, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ObsService {
  obs!:Observable<any>
  baseURL: string = "http://localhost:9000/courses";
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
  }
  validateUser(email:string,password:string):Observable<any>{
    return this.http.post<any>("http://localhost:9000/login",{"Email":email,"Password":password})
  }
  validateToken(token:any){
    return this.http.post<any>("http://localhost:9000/validateToken",{"Token":token})
  }
  getHobbiesForUser(id:string){
   return this.http.get("http://localhost:9000/users/"+id+"/hobbies") 
  }
  getCourseAndPassportForUser(id:string){
    return this.http.get("http://localhost:9000/users/"+id) 
   }
  deleteHobbyForUser(id:string){ 
    let currentToken:any=localStorage.getItem('Token')
    let headers= new HttpHeaders().set('Token',currentToken);
    return this.http.delete("http://localhost:9000/hobbies/"+id,{headers:headers}) 
  }
  createUser(email:string,password:string,firstName:string,lastName:string,address:string){
    return this.http.post<any>("http://localhost:9000/users",{"Email":email,"Password":password,"FirstName":firstName,"LastName":lastName,"Address":address})
  }
  getAllCourses(){
    return this.http.get("http://localhost:9000/courses")
  }
  enrollUserCourse(id:any,courseID:any){
    let currentToken:any=localStorage.getItem('Token')
     let headers= new HttpHeaders().set('Token',currentToken);
     console.log("Enroll course called ")
    return this.http.put<any>("http://localhost:9000/users/"+id,{"Courses":[{"ID":courseID}]},{headers:headers})
  }
  addUserHobby(id:any,hobbyName:string){
    let currentToken:any=localStorage.getItem('Token')
     let headers= new HttpHeaders().set('Token',currentToken);
    return this.http.put<any>("http://localhost:9000/users/"+id,{"Hobbies":[{"HobbyName":hobbyName}]},{headers:headers})
  }
  addPassportDetailsForUser(id:any,passportID:any,expiryDate:any){
     let currentToken:any=localStorage.getItem('Token')
     let headers= new HttpHeaders().set('Token',currentToken);
     console.log("Passport id ",Number(passportID))
    return this.http.post<any>("http://localhost:9000/users/"+id+"/passport",{"Passport":{"PassportID":Number(passportID),"ExpiryDate":expiryDate}},{headers:headers})
  }
  deleteCourseForUser(id:any,courseid:any){
     let currentToken:any=localStorage.getItem('Token')
     let headers= new HttpHeaders().set('Token',currentToken);
    return this.http.delete("http://localhost:9000/users/"+id+"/courses/"+courseid,{headers:headers}) 
  }
  addCourse(courseName:string){
    let currentToken:any=localStorage.getItem('Token')
    let headers= new HttpHeaders().set('Token',currentToken);
    return this.http.post<any>("http://localhost:9000/courses",{"Name":courseName},{headers:headers}) 
  }
}

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
    let currentToken=localStorage.getItem('Token')
    let headers:any ={}
    headers['Token']=currentToken
    console.log("Header data ",headers['Token'])
    return this.http.delete("http://localhost:9000/hobbies/"+id,{headers:headers}) 
  }
  createUser(email:string,password:string,firstName:string,lastName:string,address:string){
    return this.http.post<any>("http://localhost:9000/users",{"Email":email,"Password":password,"FirstName":firstName,"LastName":lastName,"Address":address})
  }
}

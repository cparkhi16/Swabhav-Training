import { Injectable } from '@angular/core';
import { interval, map, Observable } from 'rxjs';
import {HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ServiceService {

  constructor(private http:HttpClient) { }

  getRandomInt(max:number) {
    return Math.floor(Math.random() * max);
  }

  getData():Observable<any>{
    let obs=interval(1000).pipe(map(n=>{
      n=this.getRandomInt(36)
      var today = new Date()
      var time = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();
      return {"Data":n,"Time":time}
    }))
    return obs
  }

  getCourses():Observable<any>{
    return this.http.get<any>("http://localhost:8000/courses/");
  }

  getCourses2():Observable<any>{
    return interval(2000).pipe(map(n=>{
      return {"data":this.http.get<any>("http://localhost:8000/courses/"),"time":new Date()}
    }))
  }

  login(email:string,password:string):Observable<any>{
    return this.http.post<any>("http://localhost:8000/login",{"email":email,"password":password});
  }

  checkToken(token:string):Observable<any>{
    return this.http.post<any>("http://localhost:8000/checkToken",{"token":token});
  }
}

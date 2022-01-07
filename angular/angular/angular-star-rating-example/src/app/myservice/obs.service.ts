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
  //  console.log("My data ",d)
  //   d.pipe(map(n=>{
  //     console.log("Mydata from api ",n)
  //   }))
  //   return this.http.get(this.baseURL).pipe(map(n=>{
  //        var today = new Date()
  //     var time = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();
  //     return {"Data":n,"Time":time}
  //   }))
  // }
  }
}

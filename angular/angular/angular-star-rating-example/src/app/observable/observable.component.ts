import { Component, OnInit } from '@angular/core';
import { interval, Observable } from 'rxjs';
import { map } from 'rxjs';
import { from } from 'rxjs';
import { ObsService } from '../myservice/obs.service';

@Component({
  selector: 'app-observable',
  templateUrl: './observable.component.html',
  styleUrls: ['./observable.component.css']
})
export class ObservableComponent implements OnInit {
  myobs!:Observable<any>
  newObs!:Observable<any>
  constructor(private obs:ObsService) { }
  courses:any[]=[]
  courseName:any[]=[]
  randomData:any
  ngOnInit(): void {
    // this.getAsync3().subscribe((data)=>{
    //   console.log("Data from observable ",data)
    // },(err)=>{
    //   console.log("Error ",err)
    // },()=>{
    //   console.log("Observable complete")
    // })
    //this.getAsync3()
    //this.getAsync4()
    this.myobs=this.obs.getDataFromApi()
    this.myobs.subscribe((data)=>{
      console.log(" Data ",JSON.parse(data))
      this.courses=JSON.parse(data)
      for(let c of this.courses){
        this.courseName.push(c.Name)
      }
      console.log("Course names ",this.courseName)
    })
    this.newObs=this.obs.getRandomDataFromApi()
    this.newObs.subscribe({
  
      next:data=>{
        console.log("Data random ---",data.subscribe((dt:any)=>{
          console.log(dt)
          this.randomData=JSON.parse(dt)
        }))
        //console.log(JSON.parse(data))
      },
      error:err=>{console.log(err)},
      complete:()=>{console.log("Completed")}
    })
    console.log("Final data ",this.randomData)
  }
  getAsync1():Observable<any>{
    const obs= new Observable<any>((observer)=>{
      observer.next(10)
      observer.next(80)
      //throw new Error("a error occured ")
      observer.complete()
      observer.next(100)
    })
    return obs
  }
  getAsync2(){
    const n=interval(2000)
    n.subscribe((data)=>{
      console.log("Data returned from interval observable ",data)
    })
  }
  getAsync3():Observable<any>{
    console.log("ASYNC 3")
    return interval(2000).pipe(map((n,i)=>{
      return {"Data":n,"Time":new Date(),"Index":i}
    }))
  }
  getAsync4(){
    let srcName = from(['John', 'Tom', 'Katy'])
  srcName
   .pipe(map(data => {
     return data.toUpperCase();
   }))
   .subscribe(data => console.log(data))
}

}

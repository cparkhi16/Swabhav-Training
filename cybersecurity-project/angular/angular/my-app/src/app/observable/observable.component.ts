import { Component, OnInit } from '@angular/core';
import { from, interval, map, Observable } from 'rxjs';
import { ServiceService } from '../service/service.service';

@Component({
  selector: 'app-observable',
  templateUrl: './observable.component.html',
  styleUrls: ['./observable.component.css']
})
export class ObservableComponent implements OnInit {

  obsFromService!:Observable<any>;
  courses:any[]=[];
  constructor(private myservice:ServiceService) { }

  getAsync1():Observable<any>{
    const obs= new Observable<any>((observer)=>{
      observer.next(10)
      observer.next(80)
      //throw new Error("a error occured ")
      observer.complete()
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

  ngOnInit(): void {
    this.obsFromService=this.myservice.getData();
    // this.myservice.getCourses().subscribe((data)=>{
    //   console.log(data);
    //   this.courses=data;
    //   console.log(this.courses);
    // });

    this.myservice.getCourses2().subscribe((data)=>{
      console.log(data);
      data.data.subscribe((courses:any)=>{
        this.courses=courses;
      });
    })
    
    // this.getAsync1().subscribe((data)=>{
    //   console.log("getAsync1 data-",data);
    // },(error)=>{
    //   console.log("getAsync1 error-",error);
    // },()=>{
    //   console.log("getAsync1 completed");
    // });
    //this.getAsync2();
    // this.getAsync3().subscribe((data)=>{
    //   console.log("getAsync3 data-",data);
    // });
    // this.getAsync4();
    // this.myservice.getData().subscribe((data)=>{
    //   console.log("from service-",data);
    // })
  }

}

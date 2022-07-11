import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit ,OnDestroy{
  user: {id: number, name: string};
  parameterSubscription : Subscription;

  constructor(private route: ActivatedRoute) {
    console.log("uSER COMPONENT ",this.user)
   }

  ngOnInit() {
    this.user= {
      id: this.route.snapshot.params['id'],
      name : this.route.snapshot.params['name']
    }
    this.parameterSubscription = this.route.params.
    subscribe((params:Params)=>{
      this.user= {
        id: params['id'],
        name : params['name']
      }
    })
  }

  ngOnDestroy(): void {
    console.log("Unsubscribing params obs -=-------==")
    this.parameterSubscription.unsubscribe();
  }

}

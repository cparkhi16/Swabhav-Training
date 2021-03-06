import { Component, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  @ViewChild('f') signUpForm: NgForm;
  defaultQuestion: string = 'pet'
  answer = ''
  suggestUserName() {
    const suggestedName = 'Superuser';
  }
  // onSubmit(form: NgForm){
  //   console.log("submitted",form)
  // }
  onSubmit(){
    console.log("submitted",this.signUpForm)
  }
}

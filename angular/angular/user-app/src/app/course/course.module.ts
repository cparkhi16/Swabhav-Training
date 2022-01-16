import { FormsModule } from '@angular/forms';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';

import { CourseRoutingModule } from './course-routing.module';
import { CourseComponent } from './course.component';


@NgModule({
  declarations: [
    CourseComponent
  ],
  imports: [
    CommonModule,
    CourseRoutingModule,FormsModule,ReactiveFormsModule
  ]
})
export class CourseModule { }

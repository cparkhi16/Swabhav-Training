import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AdminCoursesRoutingModule } from './admin-courses-routing.module';
import { CourseListComponent } from './course-list/course-list.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';


@NgModule({
  declarations: [
    CourseListComponent
  ],
  imports: [
    CommonModule,
    AdminCoursesRoutingModule,
    FormsModule,
    ReactiveFormsModule,
  ],
  exports:[
    CourseListComponent
  ]
})
export class AdminCoursesModule { }

import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './auth/auth.guard';
import { CourseComponent } from './course/course.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { HobbyComponent } from './hobby/hobby.component';
import { LoginComponent } from './login/login.component';
import { PassportComponent } from './passport/passport.component';
import { RegisterComponent } from './register/register.component';

const routes: Routes = [
  {path:'',redirectTo:'login', pathMatch:'full'},
  {path:'login',component:LoginComponent},
  {path:'register',component:RegisterComponent},
  {path:'dashboard/:id',component:DashboardComponent,canActivate:[AuthGuard],children:[
    {path:'hobbies',component:HobbyComponent},
    {path:'courses',component:CourseComponent},
    {path:'passport',component:PassportComponent},
    { path: 'allCourses', loadChildren: () => import('./admin-courses/admin-courses.module').then(m => m.AdminCoursesModule) },
    { path: 'files', loadChildren: () => import('./files/files.module').then(m => m.FilesModule) },
  ]},
  {path:'**',redirectTo:'login', pathMatch:'full'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

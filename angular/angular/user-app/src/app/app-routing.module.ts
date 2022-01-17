import { ProfileComponent } from './profile/profile.component';
import { Component, Host, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { UserDetailComponent } from './user-detail/user-detail.component';
import {AuthGuardService} from './auth/auth-guard.service';

const routes: Routes = [
    {path:'',component:LoginComponent,pathMatch:"full"},
    {
        path:'register',children:[
    {path:'',component:RegisterComponent,pathMatch:"full"}
]},{
    path:'login',children:[
        {path:'',component:LoginComponent,pathMatch:"full"}]},
        {path:'profile/:userId',component:ProfileComponent,canActivate:[AuthGuardService]},
        {path:'userDetail/:userId',canActivate:[AuthGuardService],
    children:[ {path:'',component:UserDetailComponent,pathMatch:"full"}]},
    { path: 'courses', loadChildren: () => import('./course/course.module').then(m => m.CourseModule) }];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

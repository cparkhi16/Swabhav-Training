import { Component, Host, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { UserDetailComponent } from './user-detail/user-detail.component';


const routes: Routes = [
    {
        path:'register',children:[
    {path:'',component:RegisterComponent,pathMatch:"full"},{path:'userDetail',component:UserDetailComponent}
]},{
    path:'login',children:[
        {path:'',component:LoginComponent,pathMatch:"full"},{path:'userDetail',component:UserDetailComponent}]}];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
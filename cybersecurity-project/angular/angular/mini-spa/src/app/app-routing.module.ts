import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AboutComponent } from './about/about.component';
import { ContactComponent } from './contact/contact.component';
import { HomeComponent } from './home/home.component';
import { InfoComponent } from './info/info.component';
import { NotfoundComponent } from './notfound/notfound.component';
import { CanActivateTeam } from './auth.guard';

const routes: Routes = [
  { path: '', redirectTo: 'contact', pathMatch: 'full' },
  {path:'home',
  children:[
    {path:'',component:HomeComponent},
    {path:'info',component:InfoComponent,pathMatch:'full'}
  ]},
  {path:'about/:id',component:AboutComponent,canActivate: [CanActivateTeam]},
  {path:'contact',component:ContactComponent},
  { path: 'banking', loadChildren: () => import('./banking/banking.module').then(m => m.BankingModule) },
  {path:'**',component:NotfoundComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

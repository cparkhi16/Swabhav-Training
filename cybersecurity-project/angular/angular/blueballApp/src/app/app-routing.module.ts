import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { BallPageComponent } from './ball-page/ball-page.component';
import { WelcomePageComponent } from './welcome-page/welcome-page.component';

const routes: Routes = [
  { path: '', redirectTo: '/home', pathMatch: 'full' },
  {
    path:"home",
    component:WelcomePageComponent,
  },
  {
    path:"play",
    component:BallPageComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

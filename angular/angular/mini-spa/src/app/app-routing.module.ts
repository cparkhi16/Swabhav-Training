import { Host, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AboutComponent } from './about/about.component';
import { ContactComponent } from './contact/contact.component';
import { FeedbackComponent } from './feedback/feedback.component';
import { HomeComponent } from './home/home.component';
import { NotFoundComponent } from './not-found/not-found.component';

const routes: Routes = [{path:'',component:HomeComponent},
{path:'home',
children:[{path:'',component:HomeComponent,pathMatch:'full'},{path:'feedback',component:FeedbackComponent}]},
{path:'contact',component:ContactComponent},
{path:'about',component:AboutComponent},
{path:'**',component:NotFoundComponent}];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

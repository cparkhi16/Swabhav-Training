import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { AppComponent } from './app.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';

import { ParentComponent } from './parent/parent.component';
import { ChildComponent } from './child/child.component';
import { ObservableComponent } from './observable/observable.component';
import { LoginComponent } from './login/login.component';
import { BankingModule } from './banking/banking.module';
import { UserDetailComponent } from './user-detail/user-detail.component';
import { AppRoutingModule } from './app-routing.module';
import { RegisterComponent } from './register/register.component';
@NgModule({
  declarations: [
    AppComponent,
    ParentComponent,
    ChildComponent,
    ObservableComponent,
    LoginComponent,
    UserDetailComponent,
    RegisterComponent
  ],
  imports: [
    BrowserModule,
    NgbModule,FormsModule, HttpClientModule,ReactiveFormsModule,BankingModule,AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

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
@NgModule({
  declarations: [
    AppComponent,
    ParentComponent,
    ChildComponent,
    ObservableComponent,
    LoginComponent
  ],
  imports: [
    BrowserModule,
    NgbModule,FormsModule, HttpClientModule,ReactiveFormsModule,BankingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

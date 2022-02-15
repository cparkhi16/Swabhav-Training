import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ClockComponent } from './clock/clock.component';
import { TestComponent } from './test/test.component';
import { AdderPipe } from './add.pipe';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HeaderComponent } from './header/header.component';
import { FooterComponent } from './footer/footer.component';
import { SpinnerComponent } from './spinner/spinner.component';
import { LoopingComponent } from './looping/looping.component';
import { GreetingComponent } from './greeting/greeting.component';
import { TwoWayComponent } from './two-way/two-way.component';
import { PaintTaskComponent } from './paint-task/paint-task.component';
import { StudentComponent } from './student/student.component';
import { TestDirective } from './directive/test.directive';
import { NgIfNotDirective } from './directive/ng-if-not.directive';
import { HooksComponent } from './hooks/hooks.component';
import { ChildComponent } from './child/child.component';
import { ObservableComponent } from './observable/observable.component';
import { HttpClientModule } from '@angular/common/http';
import { LoginComponent } from './login/login.component';
import { BankingModule } from './banking/banking.module';

@NgModule({
  declarations: [
    AppComponent,
    ClockComponent,
    TestComponent,
    AdderPipe,
    HeaderComponent,
    FooterComponent,
    SpinnerComponent,
    LoopingComponent,
    GreetingComponent,
    TwoWayComponent,
    PaintTaskComponent,
    StudentComponent,
    TestDirective,
    NgIfNotDirective,
    HooksComponent,
    ChildComponent,
    ObservableComponent,
    LoginComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    ReactiveFormsModule,
    BankingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

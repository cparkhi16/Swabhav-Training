import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { WelcomeComponent,FileSizePipe } from './welcome.component';
import { CustomComponent } from './customcomponent/customcomponent.component';
import { FooterComponent } from './footer/footer.component';
import { HeaderComponent } from './header/header.component';
import { SpinnerComponent } from './spinner/spinner.component';
import { LoopingComponent } from './looping/looping.component';
import { TwoWayComponent } from './two-way/two-way.component';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    WelcomeComponent,
    FileSizePipe,
    CustomComponent,
    FooterComponent,
    HeaderComponent,
    SpinnerComponent,
    LoopingComponent,
    TwoWayComponent
  ],
  imports: [
    BrowserModule,FormsModule
  ],
  providers: [],
  bootstrap: [WelcomeComponent]
})
export class WelcomeModule { }

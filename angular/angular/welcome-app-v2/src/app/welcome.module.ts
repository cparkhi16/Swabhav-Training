import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { WelcomeComponent,FileSizePipe } from './welcome.component';
import { CustomComponent } from './customcomponent/customcomponent.component';

@NgModule({
  declarations: [
    WelcomeComponent,
    FileSizePipe,
    CustomComponent
  ],
  imports: [
    BrowserModule
  ],
  providers: [],
  bootstrap: [WelcomeComponent]
})
export class WelcomeModule { }

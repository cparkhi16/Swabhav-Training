import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FormsModule } from '@angular/forms';
import { ProductDisplayComponent } from './product-display/product-display.component';
import { ProductListComponent } from './product-list/product-list.component';
import { TestDirective } from './directive/test.directive';

@NgModule({
  declarations: [
    AppComponent,
    ProductDisplayComponent,
    ProductListComponent,
    TestDirective
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

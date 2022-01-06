import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { DisplayProductComponent } from './display-product/display-product.component';
import { OperateProductsComponent } from './operate-products/operate-products.component';
import { IfDirective } from 'src/structural-directive/my.directive';
import { NullDefaultValueDirective } from './attribute-directive/my-directive.directive';

@NgModule({
  declarations: [
    AppComponent,
    DisplayProductComponent,
    OperateProductsComponent,
    IfDirective,
    NullDefaultValueDirective
  ],
  imports: [
    BrowserModule,FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

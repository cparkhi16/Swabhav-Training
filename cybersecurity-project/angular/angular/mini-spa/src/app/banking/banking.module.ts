import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { BankingRoutingModule } from './banking-routing.module';
import { BankingComponent } from './banking.component';
import { AccountComponent } from './account/account.component';


@NgModule({
  declarations: [
    BankingComponent,
    AccountComponent
  ],
  imports: [
    CommonModule,
    BankingRoutingModule
  ]
})
export class BankingModule { 
  constructor(){
    console.log("banking module loaded");
  }
}

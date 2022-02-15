import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AccountComponent } from './account/account.component';
import { BankingComponent } from './banking.component';

const routes: Routes = [{ path: '', component: BankingComponent },
{path:"account",component:AccountComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BankingRoutingModule { }

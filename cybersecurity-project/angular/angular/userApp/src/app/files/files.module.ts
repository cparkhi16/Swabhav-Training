import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { FilesRoutingModule } from './files-routing.module';
import { FilesComponent } from './files.component';
import { ReactiveFormsModule } from '@angular/forms';


@NgModule({
  declarations: [
    FilesComponent
  ],
  imports: [
    CommonModule,
    FilesRoutingModule,
    ReactiveFormsModule
  ]
})
export class FilesModule { }

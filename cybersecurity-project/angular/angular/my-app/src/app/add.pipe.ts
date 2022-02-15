import { 
    Pipe, 
    PipeTransform 
 } from '@angular/core';  
 
 @Pipe ({ 
    name: 'Adder' 
 }) 
 
 export class AdderPipe implements PipeTransform { 
    transform(value: number, addValue: number): number { 
       return addValue + value 
    } 
 } 
import { Directive, ElementRef } from '@angular/core'; 
@Directive({ 
   selector: '[appCustomstyle]' 
}) 
export class CustomDirective {
   constructor(el: ElementRef) { 
      //el.nativeElement.style.fontSize = '24px'; 
      el.nativeElement.style.color='blue';
   } 
}
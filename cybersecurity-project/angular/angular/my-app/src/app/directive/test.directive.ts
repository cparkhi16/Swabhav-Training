import { Directive, ElementRef } from '@angular/core';

@Directive({
  selector: '[appTest]'
})
export class TestDirective {

  constructor(private el:ElementRef) {
    console.log("directive test here",el);
    el.nativeElement.style.backgroundColor="red";
  }

}

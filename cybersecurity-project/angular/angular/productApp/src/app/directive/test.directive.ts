import { Directive, ElementRef, HostListener } from '@angular/core';
import { NgControl } from '@angular/forms';

@Directive({
  selector: 'input[appTest]'
})
export class TestDirective {

  constructor(private eleRef: ElementRef,private control: NgControl) { 
    
  }

  // @HostListener('change', ['$event']) onChange(){
  //   console.log("here")
  //   this.eleRef.nativeElement.style.color = 'red';
  //   console.log("value-",this.eleRef.nativeElement.value);
  //   if(this.eleRef.nativeElement.value===""){
  //     this.eleRef.nativeElement.value=undefined;
  //   }
  // }

  @HostListener('click', ['$event.target'])
  onEvent(target: HTMLInputElement){
    console.log("target-value",target.value);
    this.control.viewToModelUpdate((target.value === '') ? null : target.value);
  }
  @HostListener('change', ['$event.target'])
  onChange(target: HTMLInputElement){
    console.log("target-value",target.value);
    this.control.viewToModelUpdate((target.value === '') ? null : target.value);
  }
}

import { Directive, HostBinding, HostListener,ElementRef } from "@angular/core";

@Directive({
    'selector':'[appDropdown]'
})
export class DropDownDirective{
    constructor(private elRef: ElementRef){}
    @HostBinding('class.open')  isOpen = false;
    // @HostListener('click') toggleOpen(){
    //     this.isOpen = !this.isOpen;
    // }


    //Closing the Dropdown From Anywhere by placing the listener not on the dropdown but on the document
    @HostListener('document:click', ['$event']) toggleOpen(event: Event) {
        this.isOpen = this.elRef.nativeElement.contains(event.target) ? !this.isOpen : false;
      }
}
import { Directive, Input, TemplateRef, ViewContainerRef } from '@angular/core';

@Directive({
  selector: '[appNgIfNot]'
})
export class NgIfNotDirective {

  constructor(private templateRef: TemplateRef<any>,private viewContainer: ViewContainerRef) { }

  @Input() set appNgIfNot(condition: boolean) {
    if (!condition) {
      this.viewContainer.createEmbeddedView(this.templateRef);
    }else{
      this.viewContainer.clear(); 
    }
  }

}

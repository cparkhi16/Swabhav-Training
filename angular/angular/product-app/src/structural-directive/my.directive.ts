import { Directive, Input,TemplateRef,ViewContainerRef } from '@angular/core'; 


@Directive({
    selector: '[if]'
  })
  export class IfDirective {
    variable!:boolean
    constructor(
      private template: TemplateRef<any>,
      private container: ViewContainerRef
    ) { }
  
    /*@Input()
    set delayRendering(delayTime: number) { }*/

    @Input()
    set if(variable: boolean) {this.variable=variable }
    ngOnInit() {
       /* setTimeout(() => {
          this.container.createEmbeddedView(this.template);
        }, this.delayRendering);*/
        if (this.variable==false){
            this.container.createEmbeddedView(this.template);
        }
      }
  }
import { Directive, Input,TemplateRef,ViewContainerRef } from '@angular/core'; 


@Directive({
    selector: '[delayRendering]'
  })
  export class DelayRenderingDirective {
    variable!:string
    constructor(
      private template: TemplateRef<any>,
      private container: ViewContainerRef
    ) { }
  
    /*@Input()
    set delayRendering(delayTime: number) { }*/

    @Input()
    set delayRendering(variable: string) {this.variable=variable }
    ngOnInit() {
       /* setTimeout(() => {
          this.container.createEmbeddedView(this.template);
        }, this.delayRendering);*/
        if (this.variable=="false"){
            this.container.createEmbeddedView(this.template);
        }
      }
  }
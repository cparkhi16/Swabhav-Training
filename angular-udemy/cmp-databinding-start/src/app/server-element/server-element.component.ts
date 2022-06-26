import { AfterContentChecked, AfterContentInit, AfterViewChecked, AfterViewInit, Component, ContentChild, DoCheck, ElementRef, Input, OnChanges, OnDestroy, OnInit, SimpleChanges, ViewChild, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-server-element',
  templateUrl: './server-element.component.html',
  styleUrls: ['./server-element.component.css'],
  //encapsulation : ViewEncapsulation.Emulated // by default it is emulated that is view encapsulation is enabled ( css to component binding)
  //encapsulation : ViewEncapsulation.None // it disables view encapsulation
})
export class ServerElementComponent implements OnInit ,OnChanges,DoCheck,AfterContentInit,AfterContentChecked,AfterViewInit,AfterViewChecked,OnDestroy{
  @Input('srvElement') element : {type:string,name:string,content:string};
  @Input() name:string;
  @ViewChild('heading',{static:true}) heading:ElementRef;
  @ContentChild('paragraph',{static:true}) paragraph:ElementRef;
  constructor() { 
    console.log("-- constructor called --")
  }

  ngOnInit(): void {
    console.log("ng On init called ")
    console.log("heading local reference template div content in ngOnInit ",this.heading.nativeElement.textContent)
    console.log("Paragraph val in ngOnInit ",this.paragraph.nativeElement.textContent)
  }
  ngOnChanges(changes: SimpleChanges): void {
    console.log("ngDoChange called ",changes)
  }
  ngDoCheck(): void {
    console.log("ngDocheck called ")
  }
  ngAfterContentInit(): void {
    console.log("ngAftercontent init called")
  }
  ngAfterContentChecked(): void {
    console.log("ngAfterContent checked ")
    console.log("Paragraph val in ngAfterContent checked",this.paragraph.nativeElement.textContent)
  }
  ngAfterViewInit(): void {
    console.log("ngafter view init called ")
    console.log("heading local reference template div content in ngViewInit ",this.heading.nativeElement.textContent)
  }
  ngAfterViewChecked(): void {
    console.log("ngAfter view checked")
  }
  ngOnDestroy(): void {
    console.log('ngOn destroy calleds')
  }
}

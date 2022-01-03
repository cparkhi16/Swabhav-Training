import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-footer',
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.css']
})
export class FooterComponent implements OnInit {
  date: any = new Date()
  Year:any
  constructor() { 
    this.Year=this.date.getFullYear();
  }

  ngOnInit(): void {
  }

}

import { Component, OnInit } from '@angular/core';
import { take } from 'rxjs';
import { ActivatedRoute } from '@angular/router';
@Component({
  selector: 'app-notfound',
  templateUrl: './notfound.component.html',
  styleUrls: ['./notfound.component.css']
})
export class NotfoundComponent implements OnInit {
  constructor(private route: ActivatedRoute) { }

  ngOnInit(): void {
  }

}

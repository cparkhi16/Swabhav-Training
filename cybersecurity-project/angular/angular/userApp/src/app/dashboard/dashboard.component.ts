import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api/api.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  constructor(private api:ApiService) {
    this.api.getUser(this.api.decryptData(localStorage.getItem('userId')!));
  }

  ngOnInit(): void {
  }

}

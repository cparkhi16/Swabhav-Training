import { Component, OnInit } from '@angular/core';
import { ServersService } from './servers.service';

@Component({
  selector: 'app-servers',
  templateUrl: './servers.component.html',
  styleUrls: ['./servers.component.css']
})
export class ServersComponent implements OnInit {
  public servers: {id: number, name: string, status: string}[] = [];

  constructor(private serversService: ServersService) { 
    console.log('hi---------')
  }

  ngOnInit() {
    this.servers = this.serversService.getServers();
    console.log("this servers ",this.servers)
  }

}
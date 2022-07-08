import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Route, Router } from '@angular/router';
import { ServersService } from './servers.service';

@Component({
  selector: 'app-servers',
  templateUrl: './servers.component.html',
  styleUrls: ['./servers.component.css']
})
export class ServersComponent implements OnInit {
  public servers: {id: number, name: string, status: string}[] = [];

  constructor(private serversService: ServersService,private router: Router,private route: ActivatedRoute) { 
    console.log('hi---------')
  }

  ngOnInit() {
    this.servers = this.serversService.getServers();
    console.log("this servers ",this.servers)
  }

  onReload(){
    // this.router.navigate(['servers'],{relativeTo: this.route}) navigate will not know current path so we need to give the extra js object for relative path
  }

}

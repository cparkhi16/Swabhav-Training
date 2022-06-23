import { Component } from "@angular/core";

@Component({
    selector: 'app-server',
    templateUrl : './server.component.html'
})
export class ServerComponent {
    id: number=10;
    allowAddServer:boolean = true;
    status: string ="offline";
    msg= "No server was created !"
    userName = ""
    getStatus(){
        return this.status;
    }
    constructor(){
        setTimeout(()=>{
            console.log("after 5 secs")
            this.allowAddServer=false;
        },5000)
    }
    onClick(){
        this.msg= "server was created "+this.userName
    }
    onAssignment(){
        this.userName = "";
    }
    getIsUserEmpty(){
        if(this.userName == "")
        return true;
        else
        return false;
    }
}
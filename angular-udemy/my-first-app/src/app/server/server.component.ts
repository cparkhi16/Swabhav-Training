import { Component } from "@angular/core";

@Component({
    selector: 'app-server',
    templateUrl : './server.component.html',
    styles:[
        `
        .online{
                color: white;
            }
        `
    ]
})
export class ServerComponent {
    id: number=10;
    allowAddServer:boolean = true;
    status: string ="offline";
    msg= "No server was created !"
    userName = ""
    serverCreated = false;
    show = true;
    logs=[]
    getStatus(){
        return this.status;
    }
    constructor(){
        // setTimeout(()=>{
        //     console.log("after 5 secs")
        //     this.allowAddServer=false;
        // },5000)
    }
    onClick(){
        console.log(" on click of add server ")
        this.serverCreated = true;
        this.status="online"
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
    getColor(){
        return this.serverCreated === true ? 'green' : 'red';
    }
    addLogs(){
        //this.show = !this.show
        this.logs.push(this.logs.length +1 )
    }
}
const EventEmitter=require('events')

class MyEmitter extends EventEmitter{
    someChange(message){
        this.emit("someEvent",{message})
        //console.log({message})
    }
}

module.exports=MyEmitter
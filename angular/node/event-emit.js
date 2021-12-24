const EventEmitter = require('events');
class Emitter extends EventEmitter {
    someChange(){
        this.emit('event', 'a', 'b','Event emitted')
    }
}
module.exports=Emitter
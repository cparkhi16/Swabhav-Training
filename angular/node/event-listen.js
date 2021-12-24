const MyEmitter=require('./event-emit')
const eventEmitter = new MyEmitter();
eventEmitter.on('event', (a, b,c) => {
  console.log(a, b,c);
});

eventEmitter.someChange()


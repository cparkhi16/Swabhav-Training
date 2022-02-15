const MyEmitter = require('./myemitter')
//const MyEmmiter=require('./myemitter')

const myEmitter=new MyEmitter()

myEmitter.on("someEvent",function(a){
    console.log("someEvent caught",a);
})

myEmitter.on("someEvent",(a)=>{
    console.log("second event listener-",a);
})

myEmitter.someChange("hello");
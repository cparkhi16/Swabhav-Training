const server=require('http').createServer((req,res)=>{
    
})

server.on('connection', (stream) => {
    console.log('someone connected!');
    });

server.listen(3000)
console.log("server at 3000")
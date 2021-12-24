const server=require('http').createServer((req,res)=>{
  server.on('connection',(stream)=>{
      console.log("Someone connected")
  });
})

server.listen(3000)
console.log("Listening on 3000")
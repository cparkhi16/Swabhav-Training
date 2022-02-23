const express=require('express')
const cors=require('cors')
const uuid=require('uuid')
const bodyParser=require('body-parser');
const { default: axios } = require('axios');
const TodoTaskList = require('./todomodel');
const app = express();
const jwt = require("jsonwebtoken")
app.use(cors())
app.use(bodyParser.json())

const todoTasks =[]
app.post('/api/v1/generateToken',(req,resp)=>{
    //const {name}=req.body;
    console.log(req.body)
    const accessToken = generateAccessToken ({service: req.body.name})
    console.log("My access token for service ",req.body.name,accessToken)
    resp.json ({accessToken: accessToken})
})
function generateAccessToken(service) {
    console.log("Service name ",service)
    return jwt.sign(service, "chinmay", {expiresIn: "15m"}) 
}
app.get('/api/v1/tasks',(req,resp)=>{
    TodoTaskList.findAll(function(err, todoTasks) {
        console.log('controller')
        if (err)
        resp.send(err);
        console.log('res', todoTasks);
        resp.send(todoTasks);
      });
    //resp.send(todoTasks)
})
app.delete('/api/v1/tasks/:taskId',async(req,resp)=>{
    const id = req.params.taskId
    // for( var i = 0; i < todoTasks.length; i++){ 
    //     if ( todoTasks[i].id === id) { 
    //         todoTasks.splice(i, 1); 
    //     }
    // }
    await axios.post("http://eventbus-service:4005/eventbus/event",{
        type:"Remove to do task",
        data:{id}
    }).catch(e=>console.log(e.message))
    // resp.send({})

    //db

    TodoTaskList.delete( req.params.taskId, function(err, employee) {
        if (err)
        resp.send(err);
        resp.json({ error:false, message: 'To do task successfully deleted' });
      });
})
app.post('/api/v1/task',validateToken,async (req,resp)=>{
    const {task,userID}=req.body;
//     const id=uuid.v4();
//    // posts[id]={id,task}
//    todoTasks.push({id,task})
    const id=uuid.v4();
    await axios.post("http://eventbus-service:4005/eventbus/event",{
        type:"Task Created",
        data:{id,task,userid:userID}
    }).catch(e=>console.log(e.message))

    // db 
    req.body.id=id;
    const new_task = new TodoTaskList(req.body);
    console.log("Adding task for userid ",new_task.userid)
    console.log("Create todotask called ")
    //handles null error 
   if(req.body.constructor === Object && Object.keys(req.body).length === 0){
        resp.status(400).send({ error:true, message: 'Please provide all required field' });
    }else{
        TodoTaskList.create(new_task, function(err, task) {
            if (err)
            resp.send(err);
            console.log("Task id ",task)
            resp.json({error:false,message:"task added successfully!",data:task});
        });
    }
    //resp.status(201).send(todoTasks[id])
});
app.post('/eventbus/event/listener',(req,resp)=>{
    const {type}=req.body
    console.log("Received event ",type)
    resp.send({})
})
app.listen(4001,()=>{
    console.log("Todotask has started on 4001")
})
function validateToken(req, res, next) {
    //get token from request header
    console.log("Validating token ")
    const authHeader = req.headers["authorization"]
    console.log("Auth header ",authHeader)
    if(authHeader!==undefined){
    jwt.verify(authHeader, "chinmay", (err, user) => {
        console.log("Err in token ",err)
    if (err) { 
     res.status(403).send("Token invalid")
     }
     else {
     next() //proceed to the next action in the calling function
     }
    })
 } //end of jwt.verify()
 else{
    res.status(403).send("Token not present")
 }
    }
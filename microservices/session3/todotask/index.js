const express=require('express')
const cors=require('cors')
const uuid=require('uuid')
const bodyParser=require('body-parser');
const { default: axios } = require('axios');
const TodoTaskList = require('./todomodel');
const app = express();

app.use(cors())
app.use(bodyParser.json())

const todoTasks =[]

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
app.post('/api/v1/task',async (req,resp)=>{
    const {task}=req.body;
//     const id=uuid.v4();
//    // posts[id]={id,task}
//    todoTasks.push({id,task})
    const id=uuid.v4();
    await axios.post("http://eventbus-service:4005/eventbus/event",{
        type:"Task Created",
        data:{id,task}
    }).catch(e=>console.log(e.message))

    // db 
    req.body.id=id;
    const new_task = new TodoTaskList(req.body);
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
    console.log("Blogpost has started on 4001")
})
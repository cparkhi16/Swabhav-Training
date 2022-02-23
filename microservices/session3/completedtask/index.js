const express=require('express')
const cors=require('cors')
const uuid=require('uuid')
const bodyParser=require('body-parser')
const axios=require('axios')
const app = express();
const CompletedTaskList = require('./completetaskmodel');
app.use(cors())
app.use(bodyParser.json())

const completedTasks=[]
app.get('/api/v1/completed/tasks',(req,resp)=>{
    CompletedTaskList.findAll(function(err, completedTasks) {
        console.log('controller')
        if (err)
        resp.send(err);
        console.log('res', completedTasks);
        resp.send(completedTasks);
      });
    //resp.send(comments) 
})

app.post('/api/v1/completed/tasks/:userID',async (req,resp)=>{
    
    const {task}=req.body.p;
    console.log("add this to completed task ",task)
    // const {id}=req.body.p;
    const userID=req.params.userID
    // const postId=req.params.postId
    // const comment=postsWithComments[postId]||[]
    // comment.push({commentid,message})
    // postsWithComments[postId]=comment


    // completedTasks.push({id,task})
    // console.log("==== hre ",task,id,req.body.p)
    const id=uuid.v4();
    await axios.post("http://eventbus-service:4005/eventbus/event",{
        type :"Completed Task Created",
        data:{id,task,userid:userID}
    }).catch(e=>console.log(e.message))

    // resp.status(201).send({id,task})

    //db
    req.body.p.id=id;
    req.body.p.userID=userID;
    const new_task = new CompletedTaskList(req.body.p);
    console.log("Create completed task called ")
    //handles null error 
   if(req.body.constructor === Object && Object.keys(req.body).length === 0){
        resp.status(400).send({ error:true, message: 'Please provide all required field' });
    }else{
        CompletedTaskList.create(new_task, function(err, task) {
            if (err)
            resp.send(err);
            resp.json({error:false,message:"completed task added successfully!",data:task});
        });
    }
})
// app.post('/api/v1/blog/post/:postId/comment',async (req,resp)=>{
//     const commentid=uuid.v4();
//     const {message}=req.body;
//     const postId=req.params.postId
//     const comment=postsWithComments[postId]||[]
//     comment.push({commentid,message})
//     postsWithComments[postId]=comment

//     await axios.post("http://localhost:4005/eventbus/event",{
//         type :"Comment Created",
//         data:{postId,commentid,message}
//     }).catch(e=>console.log(e.message))

//     resp.status(201).send({commentid,message})
// })
app.delete('/api/v1/completed/tasks/:taskId',async(req,resp)=>{
    const id = req.params.taskId
    // for( var i = 0; i < completedTasks.length; i++){ 
    //     if ( completedTasks[i].id === id) { 
    //         completedTasks.splice(i, 1); 
    //     }
    // }
    await axios.post("http://eventbus-service:4005/eventbus/event",{
        type:"Remove completed task",
        data:{id}
    }).catch(e=>console.log(e.message))

    CompletedTaskList.delete( req.params.taskId, function(err, employee) {
        if (err)
        resp.send(err);
        resp.json({ error:false, message: 'Completed task successfully deleted' });
      });

})
app.post('/eventbus/event/listener',(req,resp)=>{
    const {type}=req.body
    console.log("Received event ",type)
    resp.send({})
})
app.listen(4002,()=>{
    console.log("completedtask has started at 4002")
});
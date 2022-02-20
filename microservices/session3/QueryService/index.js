const express=require('express')
const axios=require('axios')
const bodyParser=require('body-parser')
const cors=require('cors')

const app=express();
app.use(bodyParser.json());
app.use(cors())

const todoTasks=[]
const completedTasks=[]
app.get('/api/v1/tasks',(req,resp)=>{
    console.log("Get called ",todoTasks)
    resp.send(todoTasks);
})
app.get('/api/v1/completed/tasks',(req,resp)=>{
    console.log("Get called for completed tasks ",completedTasks)
    resp.send(completedTasks);
})
const handleMyEvent=(type,data)=>{
    console.log("type ",type,data)
    //console.log("Type of data ",typeof(data))
    if(type=="Task Created"){
        //const {id,task}=data;
        if(typeof(data)=='string'){
            data=JSON.parse(data)
        }
        todoTasks.push(data)
        //console.log("To do tasks ",todoTasks)
        return;
    }
    if(type=="Completed Task Created"){
       // const {id,task}=data
        // const post=posts[postId]
        // post.comments.push({id,message})
        if(typeof(data)=='string'){
            data=JSON.parse(data)
        }
        completedTasks.push(data)
        //console.log("completed tasks",completedTasks)
        return;
    }
    if(type == "Remove to do task"){
        if(typeof(data)=='string'){
            data=JSON.parse(data)
        }
        const {id}=data
        for( var i = 0; i < todoTasks.length; i++){ 
            if ( todoTasks[i].id === id) { 
                todoTasks.splice(i, 1); 
            }
        }
    }
    if(type == "Remove completed task"){
        if(typeof(data)=='string'){
            data=JSON.parse(data)
        }
        const {id}=data
        for( var i = 0; i < completedTasks.length; i++){ 
           // console.log("-= ",completedTasks)
            if ( completedTasks[i].id === id) { 
                completedTasks.splice(i, 1); 
            }
        }
    }
}
app.post('/eventbus/event/listener',(req,resp)=>{
    const {type,data}=req.body
    console.log("Received event",{type})
    handleMyEvent(type,data)
    resp.send({});
})

app.listen(4003,async()=>{
    const resp=await axios.get("http://localhost:4005/eventbus/event").catch(e=>console.log(e.message))
    //console.log("Resp from event bus ",resp)
    const events=resp.data || [];
    for(let e of events){
        handleMyEvent(e.type,e.data)
    }
    //console.log(resp)
    console.log("Query has loaded all the events and has started at 4003")
})

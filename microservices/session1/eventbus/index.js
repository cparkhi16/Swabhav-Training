const express=require('express')
const axios=require('axios')
const bodyParser=require('body-parser')
const cors=require('cors')

const app=express();
app.use(bodyParser.json());
app.use(cors())

const events=[]

app.get('/eventbus/event',(req,resp)=>{
    resp.send(events);
});// used by query svc when it is up to get all events

app.post('/eventbus/event',(req,resp)=>{// when a event occurs store it in events list (used by all svcs)
    const event = req.body;
    events.push(event)
    axios.post("http://blogpost_service:4001/eventbus/event/listener",event).catch(e=>console.log(e.message))//blogpost
    axios.post("http://blogcomment_service:4002/eventbus/event/listener",event).catch(e=>console.log(e.message))//blogcomment
    axios.post("http://query_service:4003/eventbus/event/listener",event).catch(e=>console.log(e.message))//query
    resp.send({})
})

app.listen(4005,()=>{
    console.log("Event bus started at 4005")
})
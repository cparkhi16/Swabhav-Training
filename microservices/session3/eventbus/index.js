const express=require('express')
const axios=require('axios')
const bodyParser=require('body-parser')
const cors=require('cors')
const EventBus = require('./eventbusmodel');
const app=express();
app.use(bodyParser.json());
app.use(cors())

const EventsList=[]

app.get('/eventbus/event',(req,resp)=>{
    
    // resp.send(events);

    //db
    EventBus.findAll(function(err, events) {
        console.log('controller')
        if (err)
        resp.send(err);
        console.log('res', events);
        // for(let event of events){
        //     EventsList.push(event)
        // }
        console.log("Events list ",events)
        resp.send(events);
      });
});// used by query svc when it is up to get all events

app.post('/eventbus/event',(req,resp)=>{// when a event occurs store it in events list (used by all svcs)
    const event = req.body;
    // console.log("Event ",event)
    // events.push(event)
    axios.post("http://todotask-service:4001/eventbus/event/listener",event).catch(e=>console.log(e.message))//blogpost
    axios.post("http://completedtask-service:4002/eventbus/event/listener",event).catch(e=>console.log(e.message))//blogcomment
    axios.post("http://query-service:4003/eventbus/event/listener",event).catch(e=>console.log(e.message))//query
    // resp.send({})
    console.log("Here ")
    const new_event = new EventBus(req.body);
    console.log("Create new event called ")
    //handles null error 
   if(req.body.constructor === Object && Object.keys(req.body).length === 0){
        resp.status(400).send({ error:true, message: 'Please provide all required field' });
    }else{
        EventBus.create(new_event, function(err, event) {
            if (err)
            resp.send(err);
            resp.json({error:false,message:"event added successfully!",data:event});
        });
    }
})

app.listen(4005,()=>{
    console.log("Event bus started at 4005")
})
const express=require('express')
const cors=require('cors')
const uuid=require('uuid')
const bodyParser=require('body-parser');
const { default: axios } = require('axios');

const app = express();

app.use(cors())
app.use(bodyParser.json())

const posts ={}

app.get('/api/v1/blog/post',(req,resp)=>{
    resp.send(posts)
})

app.post('/api/v1/blog/post',async (req,resp)=>{
    const {title}=req.body;
    const id=uuid.v4();
    posts[id]={id,title}
    await axios.post("http://eventbus-service:4005/eventbus/event",{
        type:"Post Created",
        data:{id,title}
    }).catch(e=>console.log(e.message))
    resp.status(201).send(posts[id])
});
app.post('/eventbus/event/listener',(req,resp)=>{
    const {type}=req.body
    console.log("Received event ",type)
    resp.send({})
})
app.listen(4001,()=>{
    console.log("Blogpost has started on 4001")
})
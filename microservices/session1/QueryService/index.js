const express=require('express')
const axios=require('axios')
const bodyParser=require('body-parser')
const cors=require('cors')

const app=express();
app.use(bodyParser.json());
app.use(cors())

const posts={}

app.get('/api/v1/blog/post',(req,resp)=>{
    resp.send(posts);
})
const handleMyEvent=(type,data)=>{
    console.log("type ",type)
    if(type=="Post Created"){
        const {id,title}=data;
        posts[id]={id,title,comments:[]}
        return;
    }
    if(type=="Comment Created"){
        const {postId,commentid,message}=data
        const post=posts[postId]
        post.comments.push({commentid,message})
        return;
    }
}
app.post('/eventbus/event/listener',(req,resp)=>{
    const {type,data}=req.body
    console.log("Received event",{type})
    handleMyEvent(type,data)
    resp.send({});
})

app.listen(4003,async()=>{
    const resp=await axios.get("http://eventbus_service:4005/eventbus/event").catch(e=>console.log(e.message))
    const events=resp.data || [];
    for(let e of events){
        handleMyEvent(e.type,e.data)
    }
    console.log(resp)
    console.log("Query has loaded all the events and has started at 4003")
})

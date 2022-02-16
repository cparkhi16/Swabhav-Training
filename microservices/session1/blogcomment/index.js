const express=require('express')
const cors=require('cors')
const uuid=require('uuid')
const bodyParser=require('body-parser')
const axios=require('axios')
const app = express();

app.use(cors())
app.use(bodyParser.json())

const postsWithComments={}
app.get('/api/v1/blog/post/:postId/comment',(req,resp)=>{
    const postId = req.params.postId
    const comments = postsWithComments[postId] || []
    resp.send(comments) 
})

app.post('/api/v1/blog/post/:postId/comment',async (req,resp)=>{
    const commentid=uuid.v4();
    const {message}=req.body;
    const postId=req.params.postId
    const comment=postsWithComments[postId]||[]
    comment.push({commentid,message})
    postsWithComments[postId]=comment

    await axios.post("http://localhost:4005/eventbus/event",{
        type :"Comment Created",
        data:{postId,commentid,message}
    }).catch(e=>console.log(e.message))

    resp.status(201).send({commentid,message})
})
app.post('/eventbus/event/listener',(req,resp)=>{
    const {type}=req.body
    console.log("Received event ",type)
    resp.send({})
})
app.listen(4002,()=>{
    console.log("Blogcomment has started at 4002")
});
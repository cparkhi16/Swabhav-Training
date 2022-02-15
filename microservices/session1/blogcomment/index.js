const express=require('express')
const cors=require('cors')
const uuid=require('uuid')
const bodyParser=require('body-parser')

const app = express();

app.use(cors())
app.use(bodyParser.json())

const postsWithComments={}
app.get('/api/v1/blog/post/:postId/comment',(req,resp)=>{
    const postId = req.params.postId
    const comments = postsWithComments[postId] || []
    resp.send(comments) 
})

app.post('/api/v1/blog/post/:postId/comment',(req,resp)=>{
    const commentid=uuid.v4();
    const {message}=req.body;
    const postId=req.params.postId
    const comment=postsWithComments[postId]||[]
    comment.push({commentid,message})
    postsWithComments[postId]=comment
    resp.status(201).send({commentid,message})
})

app.listen(4002,()=>{
    console.log("Blogcomment has started at 4002")
});
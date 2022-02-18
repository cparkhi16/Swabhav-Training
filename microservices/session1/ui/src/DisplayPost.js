import React,{useState,useEffect} from "react";
import axios from 'axios';
import CreateComment from "./CreateComment";
import DisplayComment from "./DisplayComment";


export default()=>{
    const [posts,updatePosts]=useState({})
    const loadPosts= async()=>{
        const resp = await axios.get('http://gposts.com/api/v1/blog/post/query').catch(e=>console.log(e.message))
        console.log(resp.data)
        updatePosts(resp.data)
    }
    useEffect(()=>{
        loadPosts();
    },[])
const cardofpost = Object.values(posts).map(p=>{
    return (
        <div className="card" style={{width:"30%",marginBottom:"20%"}}>
            <div className="card-body" key={p.id}>
                {p.title}
            </div>
            <div>
                <CreateComment postid={p.id} />
                <DisplayComment comments={p.comments}/>
            </div>
        </div>
    )
})
return (
<div className="d-flex flex-row flex-wrap justify-content-between">
    {cardofpost}
    </div>
)
}
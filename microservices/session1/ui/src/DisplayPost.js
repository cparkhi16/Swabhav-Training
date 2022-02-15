import React,{useState,useEffect} from "react";
import axios from 'axios';


export default()=>{
    const [posts,updatePosts]=useState({})
    const loadPosts= async()=>{
        const resp = await axios.get('http://localhost:5001/api/v1/blog/post').catch(e=>console.log(e.message))
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
        </div>
    )
})
return (
<div className="d-flex flex-row flex-wrap justify-content-between">
    {cardofpost}
    </div>
)
}
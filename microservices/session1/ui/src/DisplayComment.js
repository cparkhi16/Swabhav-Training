import React,{useState,useEffect} from "react";
import axios from 'axios';

export default ({comments})=>{
    //const [comments,updatecomments]=useState([])
    // const loadComments=async()=>{
    //     const resp=await axios.get(`http://localhost:4002/api/v1/blog/post/${postid}/comment`)
    //     updatecomments(resp.data)
    // }
    // useEffect(()=>{
    //     loadComments();
    // },[])
    console.log("Comments",comments)
    const listofcomment = comments.map(c=>{
        return(
            <li key={c.commentid}>
                {c.message}
            </li>
        )
    })
    return(
        <ol>
            {listofcomment}
        </ol>
    )
}
import React,{useState} from "react";
import axios from 'axios';

export default ({postid})=>{
    const [message,updatemessage]=useState("Your comment here")
    const handleMySubmit= async(e)=>{
        e.preventDefault();
        await axios.post(`http://posts.com/api/v1/blog/post/${postid}/comment`,{message}).catch(e=>console.log(e.message))
        updatemessage('')
    }
    return(
        <form onSubmit={handleMySubmit}>
            <div>
                <label className="form-group">
                    Comment here
                </label>
                <input type="text" value={message} onChange={(e)=> updatemessage(e.target.value)} className="form-control"/>
            </div>{message}<br/> {postid}
            <button className="btn btn-primary">Submit</button>
        </form>
    )
}
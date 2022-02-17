import React,{useState} from "react";
import axios from 'axios';
export default()=>{
    const TitlePost="Title Post"
    const [title,updatedTitle]=useState("My First Post")
    const onSubmitHandler= async(e)=>{
        e.preventDefault();
       await axios.post("http://posts.com/api/v1/blog/post",{title}).catch(e=>console.log(e.message))
       updatedTitle('')
    }
    return (
        <form onSubmit={onSubmitHandler}>
            <div className="form-group">
                <label>{TitlePost}</label>
                <input type="text" className="form-control" value={title} onChange={(e)=> updatedTitle(e.target.value)}/>
            </div>
            <button className="btn btn-primary">Submit</button>
            <div>You entered : {title}  </div>
        </form>
    )
}
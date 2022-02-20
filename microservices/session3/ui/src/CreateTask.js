import React,{useState} from "react";
import axios from 'axios';
var CryptoJS = require("crypto-js");
export default()=>{
    const Task="Add Task"
    var [task,updatedTask]=useState("My First Task")
    const onSubmitHandler= async(e)=>{
        e.preventDefault();
        var ciphertext = CryptoJS.AES.encrypt(JSON.stringify(task), 'my-secret-key@123').toString();
        console.log("my encrypted task ",ciphertext)
        task=ciphertext;
       await axios.post("http://chinmay.com:4001/api/v1/task",{task}).catch(e=>console.log(e.message))
       updatedTask('')
    }
    return (
        <form onSubmit={onSubmitHandler}>
            <div className="col-lg-8">
            <div class="card mb-4">
            <div class="card-body">
                <h3 className="text-dark">{Task}</h3>
                <input type="text" className="form-control" value={task} onChange={(e)=> updatedTask(e.target.value)}/>
            </div>
            <button className="btn btn-primary">Submit</button>
            </div></div>
            {/* <div>You entered : {task}  </div> */}
        </form>
    )
}
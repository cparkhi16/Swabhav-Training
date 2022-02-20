import React,{useState,useEffect} from "react";
import axios from 'axios';
import CreateCompletedTask from "./CreateCompletedTask";
import DisplayComment from "./DisplayCompletedTask";
import DisplayCompletedTask from "./DisplayCompletedTask";
var CryptoJS = require("crypto-js");

export default()=>{
    const [posts,updatePosts]=useState([])
    let tasks=[];
    const loadTasks= async()=>{
        const resp = await axios.get('http://localhost:4003/api/v1/tasks').catch(e=>console.log(e.message))
        // tasks.push(resp.data[0])
        // console.log(resp.data[0].id)
        // console.log(resp.data[0].task)
        console.log("To do tasks ",resp.data)
        for(let et of resp.data)
        {   
            //console.log("here ---------- ",typeof(et))
           // let t=JSON.parse(et)
            //console.log("=-= ",typeof(t))
            let id=et.id;
            let task=et.task;
            //console.log("decrypting to do task ",t[id])
            var bytes = CryptoJS.AES.decrypt(task, 'my-secret-key@123');
            var decryptedData = JSON.parse(bytes.toString(CryptoJS.enc.Utf8));
            console.log("decrypted task ",decryptedData)
            task=decryptedData;
            tasks.push({id,task})
        }
        console.log("resp task data from query ",tasks)
        
        console.log("task to do list data from query ",resp.data)
        updatePosts(tasks)
    }
    useEffect(()=>{
        loadTasks();
    },[])
    const addToCompleteTask= async(p)=>{
        var ciphertext = CryptoJS.AES.encrypt(JSON.stringify(p.task), 'my-secret-key@123').toString();
        console.log("my encrypted completed task ",ciphertext)
        p.task=ciphertext;
        await axios.post("http://localhost:4002/api/v1/completed/tasks",{p}).catch(e=>console.log(e.message))
        await axios.delete(`http://localhost:4001/api/v1/tasks/${p.id}`).catch(e=>console.log(e.message))
        loadTasks();
        console.log("I am called ",p)
    }
const cardofpost = Object.values(posts).map(p=>{
    return (
        // <form onSubmit={addToCompleteTask}>
        <div className="card" style={{width:"30%",marginBottom:"20%"}}>
            <div className="p-3 mb-2 bg-dark text-white" key={p.id}>
                {p.task}
                {/* <button className="btn btn-primary" onClick={addToCompleteTask}>Add to Completed List</button> */}
            </div>
            <div>
            <button onClick={()=>addToCompleteTask(p)} >Add To Completed Tasks</button> 
                {/* <CreateCompletedTask task={p.task} /> */}
                {/* <DisplayCompletedTask taskInfo={p.task}/> */}
            </div>
        </div>
        
        // </form>
    )
})
return (
<div className="d-flex flex-row flex-wrap justify-content-between">
    {cardofpost}
    </div>
)
}
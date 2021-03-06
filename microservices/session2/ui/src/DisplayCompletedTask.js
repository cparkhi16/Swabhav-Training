import React,{useState,useEffect} from "react";
import axios from 'axios';
var CryptoJS = require("crypto-js");
export default ()=>{
    let tasks=[]
    const [completedTasks,updatecompletedTasks]=useState([])
    const loadCompletedTasks=async()=>{
        const resp=await axios.get(`http://localhost:4003/api/v1/completed/tasks`)
        console.log("Completed tasks ",resp.data)
        for(let et of resp.data)
        {   
            console.log("here ---------- ",typeof(et))
            let id=et.id;
            let task=et.task;
            var bytes = CryptoJS.AES.decrypt(task, 'my-secret-key@123');
            var decryptedData = JSON.parse(bytes.toString(CryptoJS.enc.Utf8));
            console.log("decrypted completed task ",decryptedData)
            task=decryptedData;
            tasks.push({id,task})
        }
        updatecompletedTasks(tasks)
    }
    useEffect(()=>{
        loadCompletedTasks();
    },[])
    console.log("Task",completedTasks)
    // const listofcomment = completedTasks.map(c=>{
    //     return(
    //         <li key={c.id}>
    //             {c.task}
    //         </li>
    //     )
    // })
    
    // return(
    //     <ol>
    //         {listofcomment}
    //     </ol>
    // )
    const addToDoTask= async (p)=>{
        let task= p.task
        var ciphertext = CryptoJS.AES.encrypt(JSON.stringify(p.task), 'my-secret-key@123').toString();
        console.log("my encrypted to do task ",ciphertext)
        task=ciphertext;
        await axios.post("http://localhost:4001/api/v1/task",{task}).catch(e=>console.log(e.message))
        await axios.delete(`http://localhost:4002/api/v1/completed/tasks/${p.id}`).catch(e=>console.log(e.message))
        loadCompletedTasks();
        console.log("To do task ",p)
    }
    const cardofpost = Object.values(completedTasks).map(p=>{
        return (
            // <form onSubmit={addToCompleteTask}>
            <div className="card" style={{width:"30%",marginBottom:"20%"}}>
                <div className="p-3 mb-2 bg-dark text-white" key={p.id}>
                    {p.task}
                    {/* <button className="btn btn-primary" onClick={addToCompleteTask}>Add to Completed List</button> */}
                </div>
                <div>
                <button onClick={()=>addToDoTask(p)} >Move to Remaining task</button>
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
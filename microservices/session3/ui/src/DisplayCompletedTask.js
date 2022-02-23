import React,{useState,useEffect} from "react";
import axios from 'axios';
var CryptoJS = require("crypto-js");
export default ()=>{
    let tasks=[]
    const [completedTasks,updatecompletedTasks]=useState([])
    const loadCompletedTasks=async()=>{
        let userID=localStorage.getItem("userID")
        const resp=await axios.get(`http://chinmay.com/api/v1/completed/gettasks/${userID}`)
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
        let token = localStorage.getItem("token")
        var bytes = CryptoJS.AES.decrypt(token, 'my-secret-key@123');
        var decryptedToken = JSON.parse(bytes.toString(CryptoJS.enc.Utf8));
        let config = {
            headers: {
              authorization: decryptedToken,
            }
          }
        let userID= localStorage.getItem("userID")
        await axios.post(`http://chinmay.com/api/v1/task`,{task,userID},config).catch(e=>console.log(e.message))
        await axios.delete(`http://chinmay.com/api/v1/completed/tasks/${p.id}`).catch(e=>console.log(e.message))
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
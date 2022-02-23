import React,{useState,useEffect} from "react";
import axios from 'axios';
var CryptoJS = require("crypto-js");
export default()=>{
    const Task="Add Task"
    var [task,updatedTask]=useState("My First Task")
    const getJWT=async()=>{
        let name="ui";
        const resp= await axios.post("http://chinmay.com/api/v1/generateToken",{name}).catch(e=>console.log(e.message))
        console.log("Generated token by todotask ",resp.data)
        localStorage.setItem("token",CryptoJS.AES.encrypt(JSON.stringify(resp.data.accessToken), 'my-secret-key@123').toString())
    }
    useEffect(()=>{
        getJWT();
    },[])
    const onSubmitHandler= async(e)=>{
        e.preventDefault();
        var ciphertext = CryptoJS.AES.encrypt(JSON.stringify(task), 'my-secret-key@123').toString();
        console.log("my encrypted task ",ciphertext)
        task=ciphertext;
        let token = localStorage.getItem("token")
        var bytes = CryptoJS.AES.decrypt(token, 'my-secret-key@123');
        var decryptedToken = JSON.parse(bytes.toString(CryptoJS.enc.Utf8));
        let config = {
            headers: {
              authorization: decryptedToken,
            }
          }
        let userID = localStorage.getItem("userID");
       await axios.post("http://chinmay.com/api/v1/task",{task,userID},config).catch(e=>console.log("Enrror posting task to api ",e.message))
       updatedTask('')
    }
    return (
        <form onSubmit={onSubmitHandler}>
            <div className="col-lg-8">
            <div className="card mb-4">
            <div className="card-body">
                <h3 className="text-dark">{Task}</h3>
                <input type="text" className="form-control" value={task} onChange={(e)=> updatedTask(e.target.value)}/>
            </div>
            <button className="btn btn-primary">Submit</button>
            </div></div>
            {/* <div>You entered : {task}  </div> */}
        </form>
    )
}
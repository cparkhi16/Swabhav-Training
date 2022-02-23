// import React from "react";
// import CreateTask from "./CreateTask";
// import DisplayCompletedTask from "./DisplayCompletedTask";
// import DisplayTask from "./DisplayTask";

// export default ()=>{
//     return (
//         <div className="container p-3 mb-2 bg-secondary text-white">
//         <div>
//             <CreateTask></CreateTask>
//         </div>
//         <div className="col-lg-8">
//             <div className="card mb-4">
//             <div className="card-body">
//         <h3 className="p-3 mb-2 bg-danger text-white">Your To-do Tasks</h3>
//         <div>
//             <DisplayTask></DisplayTask>
//         </div>
//         </div></div></div>
//         <div className="col-lg-8">
//             <div className="card mb-4">
//             <div className="card-body"></div>
//         <h3 className="p-3 mb-2 bg-success text-white">Your Completed Tasks</h3>
//         <div>
//             <DisplayCompletedTask></DisplayCompletedTask>
//         </div></div></div>
//         </div>
//     )
// }

import CreateTask from "./CreateTask";
import DisplayCompletedTask from "./DisplayCompletedTask";
import DisplayTask from "./DisplayTask";
import axios from 'axios';


import React, { useState } from "react";
import ReactDOM from "react-dom";
import "./styles.css";

export default ()=>{
  // React States
  const [errorMessages, setErrorMessages] = useState({});
  const [isSubmitted, setIsSubmitted] = useState(false);


  const errors = {
    uname: "invalid username",
    pass: "invalid password"
  };

  const handleSubmit = async(event) => {
    //Prevent page reload
    event.preventDefault();

    var { uname, pass } = document.forms[0];
    ///api/v1/login
    console.log("username and pass enetered by user ",uname.value,pass.value)
    let username=uname.value;
    let password=pass.value;
    const resp =  await axios.post("http://chinmay.com/api/v1/login",{username,password}).catch(e=>console.log(e.message))
    console.log(resp.data);
   // console.log("Database value ",database)
    // Find user login info
   //const userData = database.find((user) => user.username === uname.value);

    // Compare user info
    if (resp.data.validUser) {
        console.log("User id ",resp.data.userData.id)
        localStorage.setItem("userID",resp.data.userData.id)
        setIsSubmitted(true);
    } else {
      // Username not found
    //   setErrorMessages({ name: "uname", message: errors.uname });
    alert("Invalid credentials")
    }
  };

  // Generate JSX code for error message
  const renderErrorMessage = (name) =>
    name === errorMessages.name && (
      <div className="error">{errorMessages.message}</div>
    );

  // JSX code for login form
  const renderForm = (
    <div className="form">
      <form onSubmit={handleSubmit}>
        <div className="input-container">
          <label>Username </label>
          <input type="text" name="uname" required />
          {renderErrorMessage("uname")}
        </div>
        <div className="input-container">
          <label>Password </label>
          <input type="password" name="pass" required />
          {renderErrorMessage("pass")}
        </div>
        <div className="button-container">
          <input type="submit" />
        </div>
      </form>
    </div>
  );
const showTasks=()=>{
 return (
    <div className="container p-3 mb-2 bg-secondary text-white">
    <div>
        <CreateTask></CreateTask>
    </div>
    <div className="col-lg-8">
        <div class="card mb-4">
        <div class="card-body">
    <h3 className="p-3 mb-2 bg-danger text-white">Your To-do Tasks</h3>
    <div>
        <DisplayTask></DisplayTask>
    </div>
    </div></div></div>
    <div className="col-lg-8">
        <div class="card mb-4">
        <div class="card-body"></div>
    <h3 className="p-3 mb-2 bg-success text-white">Your Completed Tasks</h3>
    <div>
        <DisplayCompletedTask></DisplayCompletedTask>
    </div></div></div>
    </div>
 )
}
  return (
    // <div className="app">
    //   <div className="login-form">
    //     <div className="title">Sighn In</div>
      <div className="container">  {isSubmitted ?  showTasks(): renderForm}
     </div>
    // </div>
       
  );
}
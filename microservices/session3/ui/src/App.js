import React from "react";
import CreateTask from "./CreateTask";
import DisplayCompletedTask from "./DisplayCompletedTask";
import DisplayTask from "./DisplayTask";

export default ()=>{
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
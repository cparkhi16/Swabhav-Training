import React from "react";
import Createpost from "./Createpost";
import DisplayPost from "./DisplayPost";

export default ()=>{
    return (
        <div className="container">
        <div>
            <Createpost></Createpost>
        </div>
        <div>
            <DisplayPost></DisplayPost>
        </div>
        </div>
    )
}
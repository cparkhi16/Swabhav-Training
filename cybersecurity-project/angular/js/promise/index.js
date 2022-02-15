let userData=[
    {"name":"martin","age":23},
    {"name":"alish","age":34},
    {"name":"kesha","age":22}
]

let hobbyData=[
    {"name":"martin","hobby":"sleeping"},
    {"name":"alish","hobby":"eating"},
    {"name":"kesha","hobby":"nothing"}
]

let courseData=[
    {"name":"martin","course":"java"},
    {"name":"alish","course":"cpp"},
    //{"name":"kesha","course":"go"}
]

function onLoad(data, successCallback, errCallback){
    if(data.length==3){
        successCallback(data);
    }
    else{
        errCallback(data);
    }
}
//Callback hell
onLoad(userData,(data)=>{
    console.log("loaded userData");
    onLoad(hobbyData,(data)=>{
        console.log("loaded hobbyData");
        onLoad(courseData,(data)=>{
            console.log("loaded courseData");
        },(err)=>{
            console.log("err in userData",err);
        })
    },(err)=>{
        console.log("err in hobbyData",err);
    });
},(err)=>{
    console.log("err in courseData",err);
})

//to avoid callback hell we use promises,
function onLoadPromise(data){
    return new Promise((resolve,reject)=>{
        //dummy condition to simulate failed api-call
        if(data.length==3){
            resolve(data);
        }
        else{
            reject(data);
        }
    });
}

//so any error in then onload block will get caught in last catch block, we don't 
onLoadPromise(userData)
.then((data)=>{
    console.log("loaded userdata");
    return onLoadPromise(hobbyData);
})
.then((data)=>{
    console.log("loaded hobbyData");
    return onLoadPromise(courseData);
})
.then((data)=>{
    console.log("loaded courseData");
})
.catch((err)=>{
    console.log("err in loading",err);
}) //wihtout catch block we get error --> uncaught promise if any promise returns an error and it is not caught
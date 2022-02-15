//Using ASYNC-AWAIT

let userData=[
    {"name":"martin","age":23},
    {"name":"alish","age":34},
    {"name":"kesha","age":22}
]

let hobbyData=[
    {"name":"martin","hobby":"sleeping"},
    {"name":"alish","hobby":"eating"},
    //{"name":"kesha","hobby":"nothing"}
]

let courseData=[
    {"name":"martin","course":"java"},
    {"name":"alish","course":"cpp"},
    {"name":"kesha","course":"go"}
]

//to avoid callback hell we use can also use async await,
function onUserLoadPromise(){
    return new Promise((resolve,reject)=>{
        //dummy condition to simulate failed api-call
        if(userData.length==3){
            resolve(userData);
        }
        else{
            reject("could not load userData");
        }
    });
}

function onHobbyLoadPromise(){
    return new Promise((resolve,reject)=>{
        const Http = new XMLHttpRequest();
        const url='https://jsonplaceholder.typicode.com/posts';
        Http.open("GET", url);
        Http.send();

        Http.onreadystatechange = (e) => {
            //console.log(Http.responseText,Http.err);
            if(this.readyState==4 && this.status!==200){
                reject("could not load hobbyData");
            }
            else{
                //console.log(Http.responseText,Http.err);
                resolve("got hobby data");
            }
        }
    });
}

function onCourseLoadPromise(){
    return new Promise((resolve,reject)=>{
        //dummy condition to simulate failed api-call
        if(courseData.length==3){
            resolve(courseData);
        }
        else{
            reject("could not load courseData");
        }
    });
}

async function test(){
    try{
        await onUserLoadPromise().then((data)=>{console.log(data);});
        await onHobbyLoadPromise().then((data)=>{console.log(data);});
        await onCourseLoadPromise().then((data)=>{console.log(data);});
    }catch(err){
        console.log(err);
    }
}

test();
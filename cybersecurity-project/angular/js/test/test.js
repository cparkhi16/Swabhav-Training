function ChangeCss() {
    document.getElementById("demo").style.fontSize = "25px"; 
    document.getElementById("demo").style.color = "red";
    document.getElementById("demo").style.backgroundColor = "yellow";        
}

function takeData() {
    let name = document.forms["myForm"]["firstName"].value;
    let city=document.forms["myForm"]["city"].value;
    let data={"name":name,"city":city}
    let datastr=JSON.stringify(data);
    console.log(datastr);
    console.log(JSON.parse(datastr));
}

document.getElementById("title").addEventListener("mouseover",()=>{
    document.getElementById("title").style.color = "red";
    document.getElementById("title").style.fontSize = "35px"; 
})

document.getElementById("title").addEventListener("mouseout",()=>{
    document.getElementById("title").style.color = "black";
    document.getElementById("title").style.fontSize = "12px"; 
})

let timeoutId=null;

document.getElementById("setTimeoutTest").addEventListener("mouseover",()=>{
    timeoutId=setTimeout(()=>{
        document.getElementById("setTimeoutTest").style.color = "red";
    },3000);
    console.log("Timer ID: " + timeoutId);
})

function changeNow(){
    clearTimeout(timeoutId);
    document.getElementById("setTimeoutTest").style.color = "blue";
}

let count=0;
var timeout2
var flag=false;

function incrementCount(){
    if (flag){
    document.getElementById("sec").innerHTML=count;
    count=count+1;
    }
}

function startCounter(){
    flag=true;
    timeout2=setInterval(incrementCount,1000);
}

function stopCounter(){
    clearInterval(timeout2)
    flag=false;
}













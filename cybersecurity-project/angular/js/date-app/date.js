
var a;
var time;
 
function increment(){
    a = new Date();
    time = a.getHours() + ':' + a.getMinutes() + ':' + a.getSeconds();
    document.getElementById("dateTime").innerHTML=time;
    count=count+1;
    console.log(count);
}

setInterval(increment,1000);


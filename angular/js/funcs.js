t=setTimeout(() => {console.log("this is the first message")}, 5000);
//clearTimeout(t);
console.log("next line after setTimeout")//Displayed first
function displayHello(){
    console.log("Hello")
}
ti=setInterval(displayHello, 1000);
console.log("next line after setInterval")//Displayed second
clearInterval(ti)
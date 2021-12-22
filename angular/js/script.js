let button = document.getElementById("button")
const sect = document.querySelector('section');
const para = document.createElement('p');
para.textContent = 'We hope you enjoyed the day.';
para.style.color = 'white';
para.style.backgroundColor = 'black';
para.style.padding = '10px';
para.style.width = '250px';
para.style.textAlign = 'center';
sect.appendChild(para);
button.addEventListener("click", () => {
    button.innerHTML="Task added successfully"
    button.style.color="Red"
    sect.removeChild(para);
  });

function mouseoverevent()  
{  
    alert("This is testing mouseover");  
} 
localStorage.setItem('test', 1)
s="Local storage "+localStorage.getItem('test')
alert(s); // 1

sessionStorage.setItem('test', 100)
g="Session storage"+sessionStorage.getItem('test')
alert(g); // after refresh: 100




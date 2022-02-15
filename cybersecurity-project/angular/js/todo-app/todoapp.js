var pendingList=JSON.parse(localStorage.getItem("pendingList"));
var completedList=JSON.parse(localStorage.getItem("completedList"));

console.log(JSON.parse(localStorage.getItem("completedList")));

function addTask(){
    let taskName=document.getElementById("taskName").value;
    if (pendingList==null){
        pendingList=[taskName]
        return
    }
    pendingList.push(taskName);
}

function toCompleted(index){
    element=pendingList[index];
    pendingList.splice(index,1);
    if (completedList==null){
        completedList=[element]
        return
    }
    completedList.push(element);
}

function toPending(index){
    element=completedList[index];
    completedList.splice(index,1);
    if (pendingList==null){
        pendingList=[element]
        return
    }
    pendingList.push(element);
}

function updateView(){
    var str = '<ol>'
    var str2='<ol>'
    var index1=0;
    var index2=0;
    if (pendingList!=null){
        pendingList.forEach(function(task) {
        str += '<li>'+ task;
        str+='  <button class="btn btn-success" onclick="toCompleted('+index1+')">✓</button></li><br>';
        index1++;
        }); 
        str += '</ol>';
        document.getElementById("pendingList").innerHTML = str;
    }
    if (completedList!=null){   
        completedList.forEach(function(task) {
        str2 += '<li>'+ task;
        str2+='  <button class="btn btn-danger" onclick="toPending('+index2+')">✕</button></li><br>';
        index2++;
        }); 
        str2 += '</ol>';
        document.getElementById("completedList").innerHTML = str2;
    }
    console.log("here")
}

function storeData(){
    localStorage.setItem("pendingList",JSON.stringify(pendingList));
    localStorage.setItem("completedList",JSON.stringify(completedList));
}

setInterval(updateView,1000);
setInterval(storeData,3000);
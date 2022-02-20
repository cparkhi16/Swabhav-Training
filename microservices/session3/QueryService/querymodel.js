'user strict';
var dbConn = require('./config/db.config');
const uuid=require('uuid')

var Query = function(taskInfo){
    this.id=taskInfo.id;
    this.task= taskInfo.task;
   // this.last_name      = employee.last_name;
    // this.email          = employee.email;
    // this.phone          = employee.phone;
    // this.organization   = employee.organization;
    // this.designation    = employee.designation;
    // this.salary         = employee.salary;
    // this.status         = employee.status ? employee.status : 1;
    // this.created_at     = new Date();
    // this.updated_at     = new Date();
};
Query.createToDoTask = function (newTask, result) {    
    dbConn.query("INSERT INTO todolist set ?", newTask, function (err, res) {
        if(err) {
            console.log("error: ", err);
            result(err, null);
        }
        else{
            console.log(res.insertId);
            result(null, res.insertId);
        }
    });           
};

Query.findAllToDoTasks = function (result) {
    dbConn.query("Select * from todolist", function (err, res) {
        if(err) {
            console.log("error: ", err);
            result(null, err);
        }
        else{
            console.log('todolist : ', res);  
            result(null, res);
        }
    });   
};

Query.deleteToDoTask = function(id, result){
    dbConn.query("DELETE FROM todolist WHERE id = ?", [id], function (err, res) {
       if(err) {
           console.log("error: ", err);
           result(null, err);
       }
       else{
           result(null, res);
       }
   }); 
};

Query.createCompletedTask = function (newTask, result) {    
    dbConn.query("INSERT INTO completedtasks set ?", newTask, function (err, res) {
        if(err) {
            console.log("error: ", err);
            result(err, null);
        }
        else{
            console.log(res.insertId);
            result(null, res.insertId);
        }
    });           
};

Query.findAllCompletedTasks = function (result) {
    dbConn.query("Select * from completedtasks", function (err, res) {
        if(err) {
            console.log("error: ", err);
            result(null, err);
        }
        else{
            console.log('todolist : ', res);  
            result(null, res);
        }
    });   
};

Query.deleteCompletedTask = function(id, result){
    dbConn.query("DELETE FROM completedtasks WHERE id = ?", [id], function (err, res) {
       if(err) {
           console.log("error: ", err);
           result(null, err);
       }
       else{
           result(null, res);
       }
   }); 
};
module.exports= Query;
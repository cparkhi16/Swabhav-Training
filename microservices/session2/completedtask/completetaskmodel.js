'user strict';
var dbConn = require('./config/db.config');
const uuid=require('uuid')

var CompletedTaskList = function(taskInfo){
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
CompletedTaskList.create = function (newTask, result) {    
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

CompletedTaskList.findAll = function (result) {
    dbConn.query("Select * from completedtasks", function (err, res) {
        if(err) {
            console.log("error: ", err);
            result(null, err);
        }
        else{
            console.log('completedtasks : ', res);  
            result(null, res);
        }
    });   
};

CompletedTaskList.delete = function(id, result){
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

module.exports= CompletedTaskList;

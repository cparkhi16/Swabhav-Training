'user strict';
var dbConn = require('./config/db.config');
const uuid=require('uuid')

var user = function(taskInfo){
    this.id=uuid.v4();
    this.username=taskInfo.id;
    this.password= taskInfo.password;
};
user.create = function (newTask, result) {    
    dbConn.query("INSERT INTO users set ?", newTask, function (err, res) {
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

user.findAll = function (result) {
    dbConn.query("Select * from users", function (err, res) {
        if(err) {
            console.log("error: ", err);
            result(null, err);
        }
        else{
            console.log('users : ', res);  
            result(null, res);
        }
    });   
};

user.delete = function(id, result){
    dbConn.query("DELETE FROM users WHERE id = ?", [id], function (err, res) {
       if(err) {
           console.log("error: ", err);
           result(null, err);
       }
       else{
           result(null, res);
       }
   }); 
};
module.exports= user;
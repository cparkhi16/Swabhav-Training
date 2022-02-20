'user strict';
var dbConn = require('./config/db.config');
const uuid=require('uuid')
/* CREATE TABLE events (
    id int NOT NULL AUTO_INCREMENT,
    type varchar(255) NOT NULL,
    data varchar(255),
    PRIMARY KEY (id)
);*/
var EventBus = function(event){
    //this.id=uuid.v4();
    this.type= event.type;
    this.data=JSON.stringify(event.data);
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

EventBus.create = function (newTask, result) {    
    dbConn.query("INSERT INTO events set ?", newTask, function (err, res) {
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

EventBus.findAll = function (result) {
    dbConn.query("Select * from events", function (err, res) {
        if(err) {
            console.log("error: ", err);
            result(null, err);
        }
        else{
            console.log('events : ', res);  
            result(null, res);
        }
    });   
};

module.exports= EventBus;
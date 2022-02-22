'user strict';

const mysql = require('mysql');
//ALTER USER 'root'@'eventbusdb-service' IDENTIFIED BY 'hello'
//local mysql db connection;
const dbConn = mysql.createConnection({
  host     : 'eventbusdb-service',
  port:3306,
  user     : 'root',
  password : 'hello',
  database : 'eventbus'
});
dbConn.connect(function(err) {
  if (err) {
    console.log(err)
  }
  console.log("Database Connected!");
});
module.exports = dbConn;
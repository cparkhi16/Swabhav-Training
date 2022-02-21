'user strict';

const mysql = require('mysql');

//local mysql db connection;
const dbConn = mysql.createConnection({
  host     : 'completedtaskdb-service',
  port:3306,
  user     : 'root',
  password : 'hello',
  database : 'teste'
});
dbConn.connect(function(err) {
  if (err) {
    console.log(err)
  }
  console.log("Database Connected!");
});
module.exports = dbConn;
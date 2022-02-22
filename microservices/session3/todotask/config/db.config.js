'user strict';

const mysql = require('mysql');
//ALTER USER 'root'@'%' IDENTIFIED WITH 'mysql_native_password' BY 'hello';
//FLUSH PRIVILEGES
//local mysql db connection
const dbConn = mysql.createConnection({
  host     : 'todolistdb-service',
  port:3306,
  user     : 'root',
  password : 'hello',
  database : 'test'
});
dbConn.connect(function(err) {
  if (err) {
    console.log(err)
  }
  console.log("Database Connected!");
});
module.exports = dbConn;
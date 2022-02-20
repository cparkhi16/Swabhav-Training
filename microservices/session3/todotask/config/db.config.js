'user strict';

const mysql = require('mysql');

//local mysql db connection
const dbConn = mysql.createConnection({
  host     : 'todolistdb-service',
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
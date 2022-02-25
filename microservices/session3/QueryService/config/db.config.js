'user strict';

const mysql = require('mysql');

//local mysql db connection
const dbConn = mysql.createConnection({
  host     : 'localhost',
  user     : 'root',
  password : 'hello',
  database : 'query'
});
dbConn.connect(function(err) {
  if (err) {
    console.log(err)
  }
  console.log("Database Connected!");
});
module.exports = dbConn;
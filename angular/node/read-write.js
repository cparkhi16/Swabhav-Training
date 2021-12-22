const fs = require("fs");

fs.writeFile('input.txt', 'Hi There !', function(err) {
   if (err) {
      return console.error(err);
   }
   console.log("Data written successfully!");
   console.log("Let's read newly written data");

   fs.readFile('input.txt', function (err, data) {
      if (err) {
         return console.error(err);
      }
      console.log("Asynchronous read: " + data.toString());
   });
});
let d= "Good afternoon !!"
fs.writeFileSync("test.txt",d)
var data = fs.readFileSync('test.txt');
console.log("Synchronous read: " + data.toString());
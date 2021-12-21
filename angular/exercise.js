//path parse relative process.arg to take arguments from cmd and show args total mem and free mem os module
const path = require('path');
   
path1 = path.parse("C:/inetpub/wwwroot/chinmay/index.html");
console.log(path1);
   
   
path2 = path.relative("C:/inetpub/wwwroot/", "C:/Users/chinmay.parkhi/OneDrive - Forcepoint/Desktop/swabhav_training/go-basics/");
console.log(path2)
path2=''
console.assert(path2,"Path.relative failed")
const myArgs = process.argv.slice(2);
console.log('myArgs: ', myArgs);


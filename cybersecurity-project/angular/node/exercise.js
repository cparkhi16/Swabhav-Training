const path=require('path')
const os=require('os')

console.log(path.parse('./file.txt'));
console.log(path.resolve())

process.argv.forEach((val, index) => {
    console.log(`${index}: ${val}`);
  });

console.log(os.freemem())
console.log(os.totalmem())

const readline = require("readline");
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

var username=""

rl.question("What is your name ? ", function(name) {
    rl.question("Where do you live ? ", function(country) {
        username=name
        console.log(`${name}, is a citizen of ${country}`);
        rl.close();
    });
});

rl.on("close", function() {
    console.log("\nclose");
    console.log(username)
    process.exit(0);
});




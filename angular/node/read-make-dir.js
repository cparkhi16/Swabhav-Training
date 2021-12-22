const fs = require('fs');
const path = require('path');
fs.readdir('./', (err, files) => {
    if (err)
      console.log(err);
    else {
      console.log("\nCurrent directory filenames:");
      files.forEach(file => {
        console.log(file);
      })
    }
  })

  filenames = fs.readdirSync('./');
  console.log("Synchronous read dir ",filenames)

   
fs.mkdir(path.join('./', 'Asynchrnous'), (err) => {
    if (err) {
        return console.error(err);
    }
    console.log('Directory created successfully!');
});

fs.mkdirSync(path.join('./', "Synchronous"));
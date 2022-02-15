const fs=require('fs')

//Readdir
fs.readdir('./',(err,data)=>{
    if (err){
        console.log(err);
    }
    console.log(data);
});

var data=fs.readdirSync('./');
console.log(data);

//Readfile
fs.readFile('./file.txt',(err,data)=>{
    if (err){
        console.log(err);
    }
    console.log(data.toString());
})
var data=fs.readFileSync('./file.txt');
console.log(data.toString());

//writefile
fs.writeFile('./file.txt',"try33",{
    encoding: "utf8",
    flag: "a",
    mode: 0o666
  },(err)=>{
    if(err){
        console.log(err);
    }
});
var data=fs.readFileSync('./file.txt');
console.log(data.toString());

fs.writeFileSync('./file.txt',"try33",{
    encoding: "utf8",
    flag: "a",
    mode: 0o666
  },(err)=>{
    if(err){
        console.log(err);
    }
});
var data=fs.readFileSync('./file.txt');
console.log(data.toString());

//Rename
fs.rename('./file.txt','file1.txt',(err) => {
    if (err) throw err;
    console.log('Rename complete!');
});


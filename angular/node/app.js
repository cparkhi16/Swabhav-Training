//console.log(global)
const t=require('./test')
function abc(){
    console.log("Called abc")
}
abc()
//t()//For running SINGLE func which is exported by test,js
t.testing()
t.coding()
console.log("app.ts --->")
console.log({exports})
console.log({require})
console.log({module})
console.log({__filename})
console.log({__dirname})

var str="Hello world"
console.log(`string is ${str}`)
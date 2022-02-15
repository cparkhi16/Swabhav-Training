const t=require('./test')
console.log({exports});
console.log("filename-",__filename);
console.log("dirname-",__dirname);
console.log("module",module)
console.log("require",require)
var str="hello"
console.log(`string is ${str}`)
function adc(){
    console.log("here")
    t.doTesting()
}

adc()
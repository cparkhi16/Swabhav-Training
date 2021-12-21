function doTesting(){
    console.log("My test")
}
function doCoding(){
    console.log("My develpment")
}
//module.exports=doTesting //for single func export
module.exports.testing=doTesting
module.exports.coding=doCoding
console.log("test.ts --->")
console.log("Exports",exports)
console.log("Require",require)
console.log("Module",module)
console.log(__filename)
console.log(__dirname)
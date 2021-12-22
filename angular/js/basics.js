// strings

let s1 = '2 + 2'              // creates a string primitive
console.log("Type of s1",typeof(s1)) //string
let s2 = new String('2 + 2')  // creates a String object
console.log("Type of s2",typeof(s2))  //object
console.log(eval(s1))         // returns the number 4
console.log(eval(s2))          // returns string '2+2'
console.log(eval(s2.valueOf()))  // returns the number 4
console.log(eval(s1.valueOf())) 
console.log("Not a string value to eval",eval(2+2))

let a='a'
let b='A'
if (a==b){
    console.log("A and B are equal") // Case dependent check
}
if (a.toUpperCase()==b.toUpperCase()){
    console.log("A and B are equal")
}

//Numbers (All numbers are considered as floating point numbers not an integer)only keeps about 17 decimal places of precision

console.log(123 === 123.0 ) // true
console.log("Using number to convert string to number",Number('123') ) // returns the number 123)
let an= '123'+4 // Concatenates and returns 1234
console.log("NAN when number cannot convert string to number ",Number("Hi"))
console.log(an)
console.log("Type of an ",typeof(an))  // String
let variable
console.log("Type of unassigned variable",typeof(variable)) // undefined
let g = 55
console.log("Type of g ",typeof(g))

//bigint manipulate primitive bigint values — which are too large to be represented by the number primitive.
console.log("Bigint type of",typeof BigInt('1') === 'bigint')  // true

//Boolean
let ab=[5]
//console.log("Boolean",Boolean(ab))
let s=[]
console.log("Type of s",typeof(s))
if ([5]===ab){
    console.log("True")
}
else{
    console.log("Not equal")
}

let sym2 = Symbol('foo')
console.log(sym2)
let sym3=Symbol('foo')
if (sym2!=sym3){
    console.log("Symbols not equal")
}
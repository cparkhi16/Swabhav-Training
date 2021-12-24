interface IPerson {
    Name: string;
    Email: string;
}

let p={Name:"C",Email:"cp@fp.com"}
let j={Name:"K",Email:"kp@fp.com"}

function printDetails(people:Array<IPerson>) {
    console.log("Using of for loop")
    for (let person of people){
        console.log(person)
        console.log("Using string interpolation ")
        console.log(`${person}`)
    }
    console.log("Using in for loop")
    for (let person in people){
        console.log(people[person])
    }
  }

printDetails([p,j])
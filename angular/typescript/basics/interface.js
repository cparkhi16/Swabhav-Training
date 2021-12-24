var p = { Name: "C", Email: "cp@fp.com" };
var j = { Name: "K", Email: "kp@fp.com" };
function printDetails(people) {
    console.log("Using of for loop");
    for (var _i = 0, people_1 = people; _i < people_1.length; _i++) {
        var person = people_1[_i];
        console.log(person);
        console.log("Using string interpolation ");
        console.log("".concat(person));
    }
    console.log("Using in for loop");
    for (var person in people) {
        console.log(people[person]);
    }
}
printDetails([p, j]);

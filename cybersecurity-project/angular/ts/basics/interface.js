var User = /** @class */ (function () {
    function User(firstName, lastName, age) {
        this.firstName = firstName;
        this.lastName = lastName;
        this.age = age;
    }
    User.prototype.getFirstName = function () {
        return this.firstName;
    };
    return User;
}());
function showDetails(people) {
    for (var _i = 0, people_1 = people; _i < people_1.length; _i++) {
        var person = people_1[_i];
        console.log(person.firstName, person.lastName, person.age);
    }
}
var shan = new User("shan", "a", 64);
console.log(shan.getFirstName());
var sumi = new User("sumi", "b");
var people = [shan, sumi];
showDetails(people);

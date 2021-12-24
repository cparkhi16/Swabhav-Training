class Student {
    constructor(age) {
        this.age = age;
    }
    get Name() {
        return this.fullName;
    }
    set Name(name) {
        this.fullName = name;
    }
}
var s = new Student(5);
console.log("Getter", s.Name);
s.Name = "Chinmay";
console.log("Getter after setting full name ", s.Name, s);
console.log(`${s}`);

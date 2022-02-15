class shape{
    private _length:number;
    public breadth:number;
    constructor(length:number,breadth:number){
        this._length=length;
        this.breadth=breadth;
    }
    get length(){
        return this._length;
    }
    set length(value){
        if(value<=0){
            console.log("length is invalid")
            this._length=1;
            return;
        }
        this._length=value;
    }
}

var shape1=new shape(2,4)
console.log(`length is ${shape1.length}`)
shape1.length=-3
console.log(`length is ${shape1.length}`)
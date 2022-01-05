import { Injectable } from '@angular/core';

class Product{
  id:number
  name:string
  desc:string
  constructor(id:number,name:string,desc:string){
    this.id=id
    this.desc=desc
    this.name=name
  }
}
@Injectable({
  providedIn: 'root'
})

export class ProductService {
  productOne=new Product(1234,"ABC","DEF")
  productTwo=new Product(5678,"PQR","XYZ")
  productThree=new Product(5688,"YUT","OLJ")
  productFour=new Product(5678,"GBV","QWE")
  productList:Array<Product>=[this.productOne,this.productTwo,this.productThree,this.productFour]
  constructor() { 
    console.log("INSTANCE OF PRODUCT SERCICE")
  }
  getRandomInt(max:number) {
    return Math.floor(Math.random() * max);
  }
  getAllProducts():Array<Product>{
    return this.productList
  }
  deleteProduct(id:number){
    console.log("In delete service ",id)
    this.productList.forEach((element,index)=>{
      console.log("Element id ",element.id)
      if(element.id==id) 
      {
        this.productList.splice(index,1);
        console.log("Splicing ")
      }
   });
   console.log(this.productList)
  }
  addProduct(id:number,desc:string,name:string){
    let newProduct = new Product(id,name,desc)
    this.productList.push(newProduct)
  }
  getMaxSellingProduct():Product{
    let t= this.getRandomInt(this.productList.length)
    return this.productList[t]
  }
}

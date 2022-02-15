import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { Product } from '../model/product';
import { ProductDisplayComponent } from '../product-display/product-display.component';

@Injectable({
  providedIn: 'root'
})
export class ProductService {
  private customSubject = new Subject<any>();
  customObservable = this.customSubject.asObservable();

  products!:Product[];

  constructor() {
    console.log("Instance of service");
    this.products=[
      {ID:2,name:"bathtub",description:"snow white color"},
      {ID:45,name:"rubber duck",description:"yellow"},
      {ID:12,name:"flower pot",description:"white and pink dots"},
      {ID:24,name:"rug",description:"brown and rough"},
      {ID:13,name:"scarf",description:"red and warm"},
      {ID:9,name:"soap",description:"organic"},
      {ID:34,name:"plate",description:"silver coating"}
    ]
  }

  getAllProducts():Product[]{
    return this.products;
  }

  addProduct(product:Product){
    this.products.push(product);
    console.log(product);
    this.customSubject.next("update");
  }

  getMaxSoldProduct():Product{
    let index=Math.floor(Math.random() * (this.products.length - 1));
    console.log(index);
    return this.products[index];
  }

  deleteProduct(ID:number){
    this.products = this.products.filter(item => item.ID !== ID);
    this.customSubject.next("update");
  }
}

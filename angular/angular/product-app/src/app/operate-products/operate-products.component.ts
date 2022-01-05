import { Component, OnInit } from '@angular/core';
import { ProductService } from '../myservice/product-service.service';

@Component({
  selector: 'app-operate-products',
  templateUrl: './operate-products.component.html',
  styleUrls: ['./operate-products.component.css']
})
export class OperateProductsComponent implements OnInit {
  show!:boolean
  products:any
  id!:number
  name!:string
  desc!:string
  productID!:number
  constructor(private service:ProductService) {
    console.log("Instance of operate-product created !",service.getAllProducts())
   }
   getProducts():void{
   console.log("Button click ")
   this.products=this.service.getAllProducts()
   console.log(this.products)
   this.show=true
   }
   addProduct():void{
     this.service.addProduct(this.id,this.desc,this.name)
   }
   deleteProduct():void{
     this.service.deleteProduct(this.productID)
   }
  ngOnInit(): void {
  }

}

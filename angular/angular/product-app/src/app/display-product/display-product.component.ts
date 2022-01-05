import { Component, OnInit } from '@angular/core';
import { ProductService } from '../myservice/product-service.service';

@Component({
  selector: 'app-display-product',
  templateUrl: './display-product.component.html',
  styleUrls: ['./display-product.component.css']
})
export class DisplayProductComponent implements OnInit {
  productName:string=""
  product:any
  constructor(private service:ProductService) { }

  ngOnInit(): void {
  }
  getBestSellingProduct():void{
    this.product=this.service.getMaxSellingProduct()
    this.productName=this.product.name
    console.log("Max selling product - ",this.product.name)
  }
}

import { Component, Input, OnInit } from '@angular/core';
import { Product } from '../model/product';
import { ProductService } from '../service/product.service';

@Component({
  selector: 'app-product-display',
  templateUrl: './product-display.component.html',
  styleUrls: ['./product-display.component.css']
})
export class ProductDisplayComponent implements OnInit {
  product!:Product;

  constructor(private productService:ProductService) {
    this.product=this.productService.getMaxSoldProduct();
  }

  updateMaxSellingProduct(){
    this.product=this.productService.getMaxSoldProduct();
  }

  ngOnInit(): void {
    this.productService.customObservable.subscribe((res) => {
      this.updateMaxSellingProduct()
    }
  );
  }

}

import { Component, EventEmitter, HostListener, OnInit, Output, ViewChild } from '@angular/core';
import { Product } from '../model/product';
import { ProductDisplayComponent } from '../product-display/product-display.component';
import { ProductService } from '../service/product.service';

@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.css']
})
export class ProductListComponent implements OnInit {
  products!:Product[];
  newID:number=0;
  newName:string="";
  newDesp:string="";
  @HostListener('click', ['$event.target'])
  onClick() {
    if(this.newName==""){
      //this.newName=undefined;
    }
 }


  constructor(private productService:ProductService) {
    this.products=this.productService.getAllProducts();
  }

  ngOnInit(): void {
  }

  deleteProduct(ID:any){
    this.productService.deleteProduct(ID);
    this.products=this.productService.getAllProducts();
  }

  addProduct(){
    let newProduct:Product={ID:this.newID,name:this.newName,description:this.newDesp};
    this.productService.addProduct(newProduct);
    this.products=this.productService.getAllProducts();
  }

}

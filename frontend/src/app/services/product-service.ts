import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';

export type Product = {
  product_id: string;
  description: string;
  amount: number;
}

@Injectable({
  providedIn: 'root',
})
export class ProductService {
  private generalUrl = 'http://localhost:5000/products';
  private specificUrl = 'http://localhost:5000/product';
  private http = inject(HttpClient);

  getAllProducts(){
    return this.http.get<Product[]>(this.generalUrl);
  }

  createProduct(prod:Product){
    return this.http.post(this.specificUrl,prod)
  }
}

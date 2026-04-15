import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { Product } from './product-service';

export type ProductToInvoice = {
  prod_to_inv_id?: number;
  invoice_id: number;
  product_id: string;
  amount: number;
}

@Injectable({
  providedIn: 'root',
})
export class ProductToInvoiceService {
  private generalUrl = 'http://localhost:5001/products-to-invoice';
  // private specificUrl = 'http://localhost:5001/invoice';
  private http = inject(HttpClient);

  getAllProductsToInvoice(){
    return this.http.get<ProductToInvoice[]>(this.generalUrl);
  }

  addProductToInvoice(body:ProductToInvoice){
    return this.http.post(this.generalUrl, body);
  }

  getProductsFromInvoiceByInvoiceId(id: number){
    return this.http.get<Product[]>(`${this.generalUrl}/invoice/${id}`)
  }

  deleteProductFromInvoice(invoice_id:number, product_id:string){
    return this.http.delete(`${this.generalUrl}/${invoice_id}/${product_id}`)
  }
}

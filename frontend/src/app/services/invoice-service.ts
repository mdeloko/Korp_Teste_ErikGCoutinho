import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';

export type Invoice = {
  invoice_id: number;
  status: string;
}

@Injectable({
  providedIn: 'root',
})
export class InvoiceService {
  private generalUrl = 'http://localhost:5001/invoices';
  private specificUrl = 'http://localhost:5001/invoice';
  private http = inject(HttpClient);


  getAllInvoices(){
    return this.http.get<Invoice[]>(this.generalUrl)
  }

  createInvoice(){
    return this.http.post(this.specificUrl,null)
  }

  updateInvoiceStatus(id:number,newStatus:string){
    return this.http.patch(`${this.specificUrl}/status/${id}`,{newStatus:newStatus})
  }
}

import { Component, inject } from '@angular/core';
import { Layout } from '../../components/layout/layout';
import { OverviewCard } from '../../components/overview-card/overview-card';
import { PrimeIcons } from 'primeng/api';
import { RecentInvoicesCard } from '../../components/recent-invoices-card/recent-invoices-card';
import { InvoiceService,Invoice } from '../../services/invoice-service';
import { ProductService } from '../../services/product-service';

@Component({
  selector: 'app-home',
  standalone:true,
  imports: [Layout, OverviewCard, RecentInvoicesCard],
  templateUrl: './home.html',
  styleUrl: './home.css',
})
export class Home {
  PrimeIcons = PrimeIcons;
  InvoiceService = inject(InvoiceService);
  ProductsService = inject(ProductService);
  totalInvoices:number =-1;
  totalOpenedInvoices:number = -1;
  totalProducts:number = -1;
  latestInvoices: Invoice[] = [];

  ngOnInit(){
    this.InvoiceService.getAllInvoices().subscribe({
      next: (invoices)=>{
        this.latestInvoices = invoices.slice(-6).reverse();
        this.totalInvoices = invoices.length;
        this.totalOpenedInvoices = invoices.filter(inv => inv.status === "opened").length
      }
    })

    this.ProductsService.getAllProducts().subscribe({
      next:(products)=>{
        this.totalProducts = products.length;
      }
    })
  }

}

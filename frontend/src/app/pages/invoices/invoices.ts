import { Component, inject } from '@angular/core';
import { Layout } from '../../components/layout/layout';
import { Button } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { SelectModule } from 'primeng/select';
import { InputNumberModule } from 'primeng/inputnumber';
import { FormsModule } from '@angular/forms';
import { PrimeIcons } from 'primeng/api';
import { InvoiceService,Invoice } from '../../services/invoice-service';
import { InvoiceList } from '../../components/invoice-list/invoice-list';
import { Product, ProductService } from '../../services/product-service';
import { ProductToInvoiceService } from '../../services/product-to-invoice-service';

@Component({
  selector: 'app-invoices',
  standalone:true,
  imports: [Layout,Button,InvoiceList,DialogModule,SelectModule,FormsModule,InputNumberModule],
  templateUrl: './invoices.html',
  styleUrl: './invoices.css',
})
export class Invoices {
  PrimeIcons = PrimeIcons;
  InvoiceService = inject(InvoiceService);
  ProductsService = inject(ProductService);
  ProductToInvoiceService = inject(ProductToInvoiceService);

  invoicesAmount:number = -1;
  invoiceList !: Invoice[];
  selectedInvoice: Invoice | null = null;
  selectedInvoiceProducts: Product[] | null = null;

  isProductModalOpen: boolean = false;
  selectedProduct: Product | null = null;
  productQuantity: number = 1;
  availableProducts: Product[] = [];

  ngOnInit(){
    this.loadInvoices();
  }

  loadInvoices() {
    this.InvoiceService.getAllInvoices().subscribe({
      next: (invoice) => {
        this.invoiceList = invoice;
        this.invoicesAmount = invoice.length;
      },
      error: (err)=> console.error(err)
    })
  }

  createInvoice() {
    this.InvoiceService.createInvoice().subscribe({
      next: (res) => {
        this.loadInvoices()
      },
      error: (err) => console.log('Erro ao criar nota: ',err)
    })
  }

  selectInvoice(invoice: Invoice) {
    this.selectedInvoice = invoice;
    this.ProductToInvoiceService.getProductsFromInvoiceByInvoiceId(this.selectedInvoice.invoice_id).subscribe({
      next: (products)=>{
        this.selectedInvoiceProducts = products
      }
    })
  }

  openProductModal() {
    this.isProductModalOpen = true;
    this.ProductsService.getAllProducts().subscribe({
      next:(products)=>{
        this.availableProducts = products
      }
    })
  }
  addProductToInvoice() {
    if(this.selectedInvoice && this.selectedProduct){
      this.ProductToInvoiceService.addProductToInvoice({
        invoice_id:this.selectedInvoice.invoice_id,
        product_id: this.selectedProduct.product_id,
        amount: this.productQuantity
      }).subscribe()
    this.isProductModalOpen = false;
    this.selectedProduct = null;
    this.productQuantity = 1;
    }
  }

  removeProductOfInvoice(){
    //@ts-ignore
    this.ProductToInvoiceService.deleteProductFromInvoice(this.selectedInvoice?.invoice_id,this.selectedProduct?.product_id).subscribe()
    //@ts-ignore
    this.ProductToInvoiceService.getProductsFromInvoiceByInvoiceId(this.selectedInvoice.invoice_id).subscribe({
      next: (products)=>{
        this.selectedInvoiceProducts = products
      }
    })
  }
  printInvoice(){
    if(this.selectedInvoice && this.selectedInvoice.status){
      this.InvoiceService.updateInvoiceStatus(this.selectedInvoice.invoice_id,'closed').subscribe({
        next:()=>{
          if(this.selectedInvoice) this.selectedInvoice.status === 'closed';
          this.loadInvoices();
        },
        error: (err) => console.error(err)
      })
    }
  }
}

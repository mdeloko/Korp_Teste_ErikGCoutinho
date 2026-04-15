import { Component, EventEmitter, Input, Output } from '@angular/core';
import { DataViewModule } from 'primeng/dataview';
import { SelectModule } from 'primeng/select';
import { TagModule } from 'primeng/tag';
import { SelectItem } from 'primeng/api';
import { Invoice } from '../../services/invoice-service';
import { FormsModule } from '@angular/forms';
import { InvoiceListItem } from '../invoice-list-item/invoice-list-item';

@Component({
  selector: 'app-invoice-list',
  standalone:true,
  imports: [DataViewModule, SelectModule, InvoiceListItem, TagModule, FormsModule],
  templateUrl: './invoice-list.html',
  styleUrl: './invoice-list.css',
})
export class InvoiceList {
  sortOptions!: SelectItem[];
  sortOrder!: number;
  sortField!: string;
  sortKey!: string;
  @Input() invoices!: Invoice[];
  @Output() onEditInvoice = new EventEmitter<Invoice>();

  ngOnInit() {
    this.sortOptions = [
        { label: 'Mais Recente Antes', value: '!invoice_id' },
        { label: 'Mais Antigo Antes', value: 'invoice_id' }
    ];
  }
  onSortChange(event: any) {
    let value = event.value;

    if (value.indexOf('!') === 0) {
        this.sortOrder = -1;
        this.sortField = value.substring(1, value.length);
    } else {
        this.sortOrder = 1;
        this.sortField = value;
    }
  }
}

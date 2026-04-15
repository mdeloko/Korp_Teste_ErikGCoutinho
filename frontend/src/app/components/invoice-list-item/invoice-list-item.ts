import { Component, EventEmitter, Input, Output } from '@angular/core';
import { TagModule } from 'primeng/tag';
import { Button } from 'primeng/button';
import { Invoice } from '../../services/invoice-service';
import { PrimeIcons } from 'primeng/api';

@Component({
  selector: 'app-invoice-list-item',
  standalone: true,
  imports: [TagModule, Button],
  templateUrl: './invoice-list-item.html',
})
export class InvoiceListItem {
  PrimeIcons = PrimeIcons;
  @Input({ required: true }) invoice!: Invoice;
  @Input() isFirst: boolean = false;

  @Output() edit = new EventEmitter<Invoice>();

  getSeverityTag() {
    return this.invoice.status === 'opened' ? 'success' : 'secondary';
  }

  getTranslatedStatus() {
    return this.invoice.status === 'opened' ? 'ABERTA' : 'FECHADA';
  }
}

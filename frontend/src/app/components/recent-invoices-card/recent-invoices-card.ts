import { Component, Input } from '@angular/core';
import { PrimeIcons } from 'primeng/api';
import { NgClass } from "@angular/common";
import { ChipIcon } from '../chip-icon/chip-icon';

@Component({
  selector: 'app-recent-invoices-card',
  standalone:true,
  imports: [NgClass,ChipIcon],
  templateUrl: './recent-invoices-card.html',
  styleUrl: './recent-invoices-card.css',
})
export class RecentInvoicesCard {
  @Input() title: string = "Title";
  @Input() opened: boolean = false;

  PrimeIcons = PrimeIcons;
  get icon(): PrimeIcons {
    return this.opened ? PrimeIcons.CLOCK : PrimeIcons.CHECK_CIRCLE;
  }

  get iconColor(): string{
    return this.opened ? "text-yellow-400 dark:text-yellow-600":"text-green-400 dark:text-green-600";
  }

  get chipLabel():string{
    return this.opened ?'Em aberto':'Fechada'
  }
}


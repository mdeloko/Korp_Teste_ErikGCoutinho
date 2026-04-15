import { Component, Input } from '@angular/core';
import { PrimeIcons } from 'primeng/api';
import { NgClass } from "@angular/common";

@Component({
  selector: 'app-overview-card',
  standalone:true,
  imports: [NgClass],
  templateUrl: './overview-card.html',
  styleUrl: './overview-card.css',
})
export class OverviewCard {
  @Input() title: string = "Title";
  @Input() icon: PrimeIcons = PrimeIcons.QUESTION_CIRCLE;
  @Input() amount: number = -1;
  @Input() iconColor: string = "text-white";
}

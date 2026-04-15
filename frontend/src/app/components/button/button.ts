import { Component, signal, Input } from '@angular/core';
import { PrimeIcons } from 'primeng/api';
import { ButtonModule, ButtonSeverity } from 'primeng/button';

@Component({
  selector: 'app-button',
  imports: [ButtonModule],
  standalone:true,
  templateUrl: './button.html',
  styleUrl: './button.css',
})
export class Button {
  @Input() label: string = "Button";
  @Input() icon: PrimeIcons = PrimeIcons.ARROW_UP_LEFT;
  @Input() isLoading: boolean = false;
  @Input() isDisabled: boolean = false;
  @Input() isRaised: boolean = false;
  @Input() severity: ButtonSeverity = "primary";
  @Input() size: "small"|"large"|undefined = undefined;


  loading = signal(this.isLoading);

  load() {
    this.loading.set(this.isLoading);
  }
}

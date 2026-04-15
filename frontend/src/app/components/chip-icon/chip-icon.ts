import { Component,Input,} from '@angular/core';
import { PrimeIcons } from 'primeng/api';
import { ChipModule } from 'primeng/chip';

@Component({
  selector: 'app-chip-icon',
  standalone:true,
  imports: [ChipModule],
  templateUrl: './chip-icon.html',
  styleUrl: './chip-icon.css',
})
export class ChipIcon {
  @Input() label:string = "Chip";
  @Input() icon:PrimeIcons = PrimeIcons.QUESTION_CIRCLE;
}

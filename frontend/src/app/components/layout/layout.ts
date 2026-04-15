import { Component, OnInit } from '@angular/core';
import { MenubarModule } from 'primeng/menubar';
import { MenuItem,PrimeIcons } from 'primeng/api';

@Component({
  selector: 'app-layout',
  standalone:true,
  imports: [MenubarModule],
  templateUrl: './layout.html',
  styleUrl: './layout.css',
})
export class Layout {
  items: MenuItem[] | undefined;
  ngOnInit(){
    this.items = [
      {
          label: 'Início',
          icon: PrimeIcons.HOME,
          routerLink: '/'
      },
      {
          label: 'Notas',
          icon: PrimeIcons.RECEIPT,
          routerLink: '/notas'
      },
      {
        label:'Produtos',
        icon: PrimeIcons.SHOPPING_BAG,
        routerLink: '/produtos',
      }
    ]
  }
}


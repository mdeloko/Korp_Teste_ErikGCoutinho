import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RecentInvoicesCard } from './recent-invoices-card';

describe('RecentInvoicesCard', () => {
  let component: RecentInvoicesCard;
  let fixture: ComponentFixture<RecentInvoicesCard>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RecentInvoicesCard],
    }).compileComponents();

    fixture = TestBed.createComponent(RecentInvoicesCard);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

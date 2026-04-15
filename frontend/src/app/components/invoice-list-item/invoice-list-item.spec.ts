import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InvoiceListItem } from './invoice-list-item';

describe('InvoiceListItem', () => {
  let component: InvoiceListItem;
  let fixture: ComponentFixture<InvoiceListItem>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [InvoiceListItem],
    }).compileComponents();

    fixture = TestBed.createComponent(InvoiceListItem);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

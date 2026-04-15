import { TestBed } from '@angular/core/testing';

import { ProductToInvoice } from './product-to-invoice-service';

describe('ProductToInvoice', () => {
  let service: ProductToInvoice;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ProductToInvoice);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

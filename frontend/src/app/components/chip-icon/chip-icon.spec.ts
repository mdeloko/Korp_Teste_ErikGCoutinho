import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChipIcon } from './chip-icon';

describe('ChipIcon', () => {
  let component: ChipIcon;
  let fixture: ComponentFixture<ChipIcon>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ChipIcon],
    }).compileComponents();

    fixture = TestBed.createComponent(ChipIcon);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

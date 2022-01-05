import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OperateProductsComponent } from './operate-products.component';

describe('OperateProductsComponent', () => {
  let component: OperateProductsComponent;
  let fixture: ComponentFixture<OperateProductsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ OperateProductsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(OperateProductsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

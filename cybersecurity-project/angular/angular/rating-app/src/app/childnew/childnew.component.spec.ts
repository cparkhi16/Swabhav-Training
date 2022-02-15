import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChildnewComponent } from './childnew.component';

describe('ChildnewComponent', () => {
  let component: ChildnewComponent;
  let fixture: ComponentFixture<ChildnewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ChildnewComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ChildnewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

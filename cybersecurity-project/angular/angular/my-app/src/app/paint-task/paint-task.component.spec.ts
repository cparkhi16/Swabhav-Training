import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PaintTaskComponent } from './paint-task.component';

describe('PaintTaskComponent', () => {
  let component: PaintTaskComponent;
  let fixture: ComponentFixture<PaintTaskComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PaintTaskComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PaintTaskComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BallPageComponent } from './ball-page.component';

describe('BallPageComponent', () => {
  let component: BallPageComponent;
  let fixture: ComponentFixture<BallPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BallPageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BallPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

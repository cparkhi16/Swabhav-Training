import { TestBed } from '@angular/core/testing';
import { WelcomeComponent } from './welcome.component';

describe('AppComponent', () => {
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        WelcomeComponent
      ],
    }).compileComponents();
  });

  it('should create the app', () => {
    const fixture = TestBed.createComponent(WelcomeComponent);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });

  it(`should have as title 'welcome-app'`, () => {
    const fixture = TestBed.createComponent(WelcomeComponent);
    const app = fixture.componentInstance;
    expect(app.title).toEqual('welcome-app');
  });

  it('should render title', () => {
    const fixture = TestBed.createComponent(WelcomeComponent);
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('.content span')?.textContent).toContain('welcome-app app is running!');
  });
});

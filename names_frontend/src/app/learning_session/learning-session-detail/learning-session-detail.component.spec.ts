import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LearningSessionDetailComponent } from './learning-session-detail.component';

describe('LearningSessionDetailComponent', () => {
  let component: LearningSessionDetailComponent;
  let fixture: ComponentFixture<LearningSessionDetailComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LearningSessionDetailComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LearningSessionDetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

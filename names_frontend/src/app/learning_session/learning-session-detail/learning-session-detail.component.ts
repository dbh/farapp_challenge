import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { LearningSession } from '../LearningSession';
import { LearningSessionService } from '../LearningSession.service';

@Component({
  selector: 'app-learning-session-detail',
  templateUrl: './learning-session-detail.component.html',
  styleUrls: ['./learning-session-detail.component.css']
})
export class LearningSessionDetailComponent implements OnInit {
  @Input() learningSession : LearningSession;
  @Output() learningSessionStatus = new EventEmitter<boolean>();
  @Output() learningSessionUpdated = new EventEmitter<LearningSession>();
 
  constructor(
    private learningSessionService: LearningSessionService) { }

  ngOnInit() {
  }

  updateLearningSession(learningSessionData: LearningSession): void {
    console.log("detail updateLearningSession learningSessionData: ",learningSessionData);
    this.learningSessionService.updateLearningSession(learningSessionData)
    .subscribe((updatedLearningSession) => {
      // let existingLearningSession = this.learningSessions.find(learningSession => learningSession.id === updatedLearningSession.id);
      // Object.assign(existingLearningSession, updatedLearningSession);
      console.log("updateLearningSession in the detail component")
      this.learningSessionUpdated.emit(learningSessionData)
    });
  }

  clearEditing(): void {
    console.log("clearEditing in the detail component")
    this.learningSessionStatus.emit(false)
  }


}

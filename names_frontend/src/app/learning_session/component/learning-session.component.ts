import { Component, Input, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { LearningSessionService } from '../LearningSession.service';

import { LearningSession } from '../LearningSession' ;
import { AuthService } from '../../auth/auth.service';
import { json } from 'body-parser';

// interface UserDetails {
//   name: string;
//   email: string;
// }

@Component({
  selector: 'app-learning-session',
  templateUrl: './learning-session.component.html',
  styleUrls: ['./learning-session.component.css']
})
export class LearningSessionComponent implements OnInit {
  learningSessions: LearningSession[];
  newLearningSession: LearningSession = new LearningSession();
  editing: Boolean = false;
  editingLearningSession: LearningSession = new LearningSession();


  profileJson: string = null;

  constructor(private learningSessionService: LearningSessionService,) { }

  ngOnInit(): void {
    this.getLearningSessions();
  
  }

  log(x) {
    console.log(`log `, x);
  }



  getLearningSessions(): void {
    this.learningSessionService.getLearningSessions().subscribe((d) => {
      console.log(d);
      this.learningSessions = d;
    });
  }

  createLearningSession(learningSessionForm: NgForm): void {
    console.log('creatLearningSession: ', learningSessionForm);
    this.learningSessionService.createLearningSession(this.newLearningSession)
      .subscribe((createLearningSession) => {
        learningSessionForm.reset();
        // this.getLearningSessions();
        // TODO Decide what to do.  Get learning sessions isn't returning a learning session when it creates one
        this.newLearningSession = new LearningSession();
        this.learningSessions.unshift(createLearningSession);
      });
  }

  deleteLearningSession(id: string): void {
    this.learningSessionService.deleteLearningSession(id)
    .subscribe((x) => {
      this.learningSessions = this.learningSessions.filter(learningSession => learningSession.id !== id);
    });
  }

  editLearningSession(learningSessionData: LearningSession): void {
    this.editing = true;
    Object.assign(this.editingLearningSession, learningSessionData);

    console.log('selected for edit ', learningSessionData);

  }

  onLearningSessionStatus(flag: boolean) {
    console.log('learning session onLearningSessionStatus ', flag);
    this.editing = false;
    this.editingLearningSession = new LearningSession();
  }

  onLearningSessionUpdated(lsVal: LearningSession) {
    console.log('onLearningSessionUpdated: ', lsVal);
    this.editingLearningSession = lsVal;

    const existingLearningSession = this.learningSessions.find(learningSession => learningSession.id === lsVal.id);
    Object.assign(existingLearningSession, lsVal);

    this.editing = false;
    this.editingLearningSession = new LearningSession();
  }
}

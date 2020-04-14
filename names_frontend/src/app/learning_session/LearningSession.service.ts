import { Injectable } from '@angular/core';
import { LearningSession } from './LearningSession';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError, Observer } from 'rxjs';
import { map, tap, catchError } from 'rxjs/operators';
import * as config from '../../assets/config/yid_config.json';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable()
export class LearningSessionService {
  private baseUrl = config.yid_api.baseUrl;
  private headers_object;

  log(x) {
    console.log(`log `, x);
  }

  constructor(private http: HttpClient) {
    this.headers_object = new HttpHeaders();
    this.headers_object.append('Content-Type', 'application/json');
    this.headers_object.append('Authorization', 'Basic ' + btoa(`${config.yid_api.user}:${config.yid_api.password}`));
   }

   getLearningSessions(): Observable<LearningSession[]> {
    console.log('getLearningSessions all');
    return this.http.get<LearningSession[]>(this.baseUrl + '/users/', httpOptions);
  }

  createLearningSession(learningSessionData: LearningSession): Observable<LearningSession> {
    return this.http.post(this.baseUrl + '/users/', learningSessionData, httpOptions)
      .pipe(
        map(response => response as LearningSession)
      );
  }

  updateLearningSession(learningSessionData: LearningSession): Observable<LearningSession> {
    return this.http.put(this.baseUrl + '/users/' + learningSessionData.id, learningSessionData, httpOptions)
      .pipe(
        map(response => response as LearningSession)
      );
  }

  deleteLearningSession(id: string): Observable<any> {
    return this.http.delete(this.baseUrl + '/users/' + id, httpOptions)
    .pipe(
      tap(_ => this.log(`deleting id=${id}`)),
      catchError(err => {
        this.log(`error deleting id=${id}`);
        return throwError(err);
      })
    );
  }

  private handleError(error: any): Promise<any> {
    console.error('Some error occured', error);
    return Promise.reject(error.message || error);
  }
}

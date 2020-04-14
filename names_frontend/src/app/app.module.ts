import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import { RouterModule, Routes } from '@angular/router';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';

import { LearningSessionService } from './learning_session/LearningSession.service';
import { LearningSessionComponent } from './learning_session/component/learning-session.component';

import { HttpClientModule } from '@angular/common/http';

import { LearningSessionDetailComponent } from './learning_session/learning-session-detail/learning-session-detail.component';

import { HeaderComponent } from './components/header/header.component';
import { FooterComponent } from './components/footer/footer.component';

import { HomeComponent } from './components/home/home.component';
import { NavbarComponent } from './components/navbar/navbar.component';

const appRoutes: Routes = [
  {
    path: 'learning-sessions',
    component: LearningSessionComponent
  },
  {
    path: '',
    // redirectTo: '/learning-sessions',
    // pathMatch: 'full'
    component: HomeComponent
  },
  { path: '**', component: PageNotFoundComponent }
];

@NgModule({
  declarations: [
    AppComponent,
    LearningSessionComponent,
    PageNotFoundComponent,
    LearningSessionDetailComponent,
    HeaderComponent,
    FooterComponent,
    HomeComponent,
    NavbarComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    HttpClientModule,
    RouterModule.forRoot(
      appRoutes,
      { enableTracing: true } // <-- debugging purposes only
    ),
    FormsModule,
    ReactiveFormsModule,
  ],
  exports: [
    RouterModule
  ],
  providers: [
      LearningSessionService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
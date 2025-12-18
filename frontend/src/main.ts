import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { AppComponent } from './app/app.component';
import { registerLocaleData } from '@angular/common';
import { en_US } from 'ng-zorro-antd/i18n';
import { provideAnimations } from '@angular/platform-browser/animations';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import localeRu from '@angular/common/locales/ru';
bootstrapApplication(AppComponent, appConfig)
  .catch((err) => console.error(err));
registerLocaleData(localeRu);
//provideAnimationsAsync();
//provideAnimations();
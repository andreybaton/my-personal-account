import { inject } from '@angular/core';
import { Routes } from '@angular/router';
import { DistlearningComponent } from './distlearning.component';
import { DistlearningService } from './distlearning.service';
import { LessonViewComponent } from './lesson-view/lesson-view.component';

export default [
    {
        path     : '',
        component: DistlearningComponent,
        resolve  : {
            // data: () => inject(DistlearningService).getData(),
        },
    },
    { 
        path: 'view', 
        component: LessonViewComponent,
    }
] as Routes;

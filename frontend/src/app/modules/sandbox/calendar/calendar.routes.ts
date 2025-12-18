import { inject } from '@angular/core';
import { Routes } from '@angular/router';
import { CalendarComponent } from './calendar.component';

export default [
    {
        path     : '',
        component: CalendarComponent,
        resolve  : {
            // data: () => inject(DistlearningService).getData(),
        },
    },
] as Routes;

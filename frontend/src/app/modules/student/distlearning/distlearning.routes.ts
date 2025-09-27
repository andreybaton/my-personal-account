import { inject } from '@angular/core';
import { Routes } from '@angular/router';
import { DistlearningComponent } from './distlearning.component';
import { DistlearningService } from './distlearning.service';


export default [
    {
        path     : '',
        component: DistlearningComponent,
        resolve  : {
            // data: () => inject(DistlearningService).getData(),
        },
    },
] as Routes;

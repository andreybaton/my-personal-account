import { inject } from '@angular/core';
import { Routes } from '@angular/router';
import { ProfileComponent } from './profile.component';



export default [
    {
        path     : '',
        component: ProfileComponent,
        resolve  : {
            // data: () => inject(DistlearningService).getData(),
        },
    },
] as Routes;

import { inject } from '@angular/core';
import { Routes } from '@angular/router';
import { SandboxChildComponent } from './sanboxChild.component';

export default [
    {
        path     : '',
        component: SandboxChildComponent,
        resolve  : {
            // data: () => inject(DistlearningService).getData(),
        },
    },
] as Routes;

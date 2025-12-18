import { inject } from '@angular/core';
import { Routes } from '@angular/router';
import { SandboxComponent } from './sandbox.component';



export default [
    {
        path     : '',
        component: SandboxComponent,
        resolve  : {
            // data: () => inject(DistlearningService).getData(),
        },
    },
] as Routes;

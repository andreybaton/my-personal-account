import { Routes } from '@angular/router';
import { LayoutComponent } from './layout/layout.component';

export const routes: Routes =[
    {
        path: '',
        component: LayoutComponent,
        data: {
            layout: 'empty'
        },
        children: [
            {path: 'distancelearning', loadChildren: () => import('./modules/student/distlearning/distlearning.routes'),},
            
        ]
    },
];

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
            {path: 'profile', loadChildren: ()=>import('./modules/student/profile/profile.routes'),},
            {path: 'sandboxButton', loadChildren: ()=>import('./modules/sandbox/sandbox-pending-button/sandbox.routes')},
            {path: 'sandboxForm', loadChildren: ()=>import('./modules/sandbox/sandbox-form/sandbox.routes')},
			{path: 'sandboxCalendar', loadChildren: ()=>import('./modules/sandbox/calendar/calendar.routes')}
        ]
    },
];

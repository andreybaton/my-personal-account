import { Component } from '@angular/core';
import { Router, RouterOutlet } from '@angular/router';

interface MenuItem {
    name: string;
    description?: string;
    route: string;
}

@Component({
    selector: 'vertical-layout',
    templateUrl: './vertical.component.html',
    imports:[RouterOutlet],
})
export class VerticalLayoutComponent {

    constructor(
        private _router: Router,
    )
    {}

    dashboards: MenuItem[] = [
        { name: 'Абитуриенту', route: '/abiturientu' },
        { name: 'Профиль', route: '/profile' },
        { name: 'Система сопрождения обучения', route: '/distancelearning' },
        { name: 'Персональные документы', route: '/persondocuments' },
        { name: 'sandboxForm', route: '/sandboxForm'},
        { name: 'sandboxButton', route:'/sandboxButton'}
    ];

    navigateTo(route: string): void {
        // Навигация по роуту
        console.log('Navigating to:', route);
        this._router.navigate([route]);
    }
    ngOnInit(): void {
         console.log('vert layout load')
    }
}
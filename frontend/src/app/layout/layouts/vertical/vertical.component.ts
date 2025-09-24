import { Component } from '@angular/core';

interface MenuItem {
    name: string;
    description?: string;
    route: string;
}

@Component({
    selector: 'vertical-layout',
    templateUrl: './vertical.component.html',

})
export class VerticalLayoutComponent {
    dashboards: MenuItem[] = [
        { name: 'Abiturientu', route: '/abiturientu' },
        { name: 'Profile', route: '/profile' },
        { name: 'Sistema soprovod', route: '/system' },
        { name: 'Moodle', route: '/moodle' },
        { name: 'Inf resources', route: '/inf' },
        { name: 'Pers dock', route: '/pers' }
    ];

    navigateTo(route: string): void {
        // Навигация по роуту
        console.log('Navigating to:', route);
        // this.router.navigate([route]);
    }
}
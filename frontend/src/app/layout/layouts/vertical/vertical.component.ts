import { Component } from '@angular/core';

interface MenuItem {
  name: string;
  description?: string;
  route: string;
}

@Component({
  selector: 'vertical-layout',
  templateUrl: './vertical.component.html',
//   styleUrls: ['./dashboard.component.scss']
})
export class VerticalLayoutComponent {
  dashboards: MenuItem[] = [
    { name: 'Project', route: '/project' },
    { name: 'Analytics', route: '/analytics' },
    { name: 'Finance', route: '/finance' },
    { name: 'Crypto', route: '/crypto' },
    { name: 'Engines', route: '/engines' },
    { name: 'EnginesAdmin', route: '/engines-admin' }
  ];

  applications: MenuItem[] = [
    { name: 'Academy', route: '/academy' },
    { name: 'Chat', route: '/chat' },
    { name: 'Contacts', route: '/contacts' },
    { name: 'ECommerce', route: '/ecommerce' }
  ];

  navigateTo(route: string): void {
    // Навигация по роуту
    console.log('Navigating to:', route);
    // this.router.navigate([route]);
  }
}
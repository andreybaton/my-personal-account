import { Component } from '@angular/core';
import { Router, RouterOutlet } from '@angular/router';

@Component({
    selector: 'horizontal-layout',
    templateUrl: './horizontal.component.html',
    imports:[],
})
export class HorizontalLayoutComponent {

  studentName = 'Иван Иванов'; // Замените на реальные данные
  studentGroup = 'ИТ-21'; // Дополнительная информация
  avatarUrl = ''; // URL аватара, если есть

  // Меню пользователя
  userMenuItems = [
    { label: 'Профиль', action: 'profile' },
    { label: 'Настройки', action: 'settings' },
    { label: 'Выйти', action: 'logout' }
  ];

  isUserMenuOpen = false;

  toggleUserMenu(): void {
    this.isUserMenuOpen = !this.isUserMenuOpen;
  }

  onMenuItemClick(action: string): void {
    this.isUserMenuOpen = false;
    console.log('Menu action:', action);
    // Здесь добавьте логику для каждого действия
  }

  ngOnInit(): void {
         console.log('hor layout load')
    }
    
}
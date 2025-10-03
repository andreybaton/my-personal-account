import { Component, OnInit } from '@angular/core';
import { ProfileService } from './profile.service';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';

interface StudentInfo {
label: string;
value: string;
icon: string;
}

@Component({
selector: 'app-distlearning',
standalone: true,
imports: [CommonModule],
templateUrl: './profile.component.html',
//   styleUrl: './distlearning.css'

})
export class ProfileComponent  {

    edit:boolean = false;

    student = {
        fullName: 'Иванов Иван Иванович',
        group: 'ИТ-21',
        email: 'student@university.ru',
        phone: '+7 (999) 123-45-67',
        faculty: 'Факультет информационных технологий',
        specialization: 'Программная инженерия',
        course: '2 курс',
        semester: '4 семестр',
        studentId: '202312345',
        admissionYear: 2023,
        avatar: '', // URL аватара, если есть
        initials: 'ИИИ' // Добавляем инициалы
    };

    // Дополнительная информация
    studentInfo: StudentInfo[] = [
        { label: 'Группа', value: this.student.group, icon: '👥' },
        { label: 'Email', value: this.student.email, icon: '📧' },
        { label: 'Телефон', value: this.student.phone, icon: '📱' },
        { label: 'Факультет', value: this.student.faculty, icon: '🏛️' },
        { label: 'Специальность', value: this.student.specialization, icon: '🎓' },
        { label: 'Курс', value: this.student.course, icon: '📚' },
        { label: 'Семестр', value: this.student.semester, icon: '📅' },
        { label: 'Студенческий билет', value: this.student.studentId, icon: '🆔' },
        { label: 'Год поступления', value: this.student.admissionYear.toString(), icon: '🎯' }
    ];

    // Статистика успеваемости
    performanceStats = {
        averageScore: 4.7,
        completedCourses: 12,
        currentCourses: 6,
        attendance: 95
    };

    // Быстрые действия
    quickActions = [
        { label: 'Редактировать профиль', icon: '✏️', action: 'edit' },
        { label: 'Скачать справку', icon: '📄', action: 'download' },
        
    ];

    constructor() {
        // Вычисляем инициалы при инициализации
        this.student.initials = this.getInitials(this.student.fullName);
    }

    // Метод для получения инициалов
    private getInitials(fullName: string): string {
        return fullName
        .split(' ')
        .map(name => name[0])
        .join('')
        .toUpperCase();
    }

    onActionClick(action: string): void {
        console.log('Action clicked:', action);
        // Здесь можно добавить логику для каждого действия
        switch (action) {
        case 'edit':
            this.editProfile();
            break;
        case 'download':
            this.downloadCertificate();
            break;
        case 'grades':
            this.viewGrades();
            break;
        case 'settings':
            this.openSettings();
            break;
        }
    }

    private editProfile(): void {
        console.log('Opening profile editor...');

        this.edit = !this.edit;
    }

    private downloadCertificate(): void {
        console.log('Downloading certificate...');
        // Логика скачивания справки
    }

    private viewGrades(): void {
        console.log('Opening grades history...');
        // Логика просмотра оценок
    }

    private openSettings(): void {
        console.log('Opening settings...');
        // Логика открытия настроек
    }
}

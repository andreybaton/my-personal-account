import { Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Lesson } from './lesson.types';
import { Router } from '@angular/router';
@Component({
    selector: 'app-lesson',
    standalone: true,
    imports: [CommonModule],
    templateUrl: './lesson.component.html',
    styles: [`
        .lesson-card {
        min-height: 60px;
        }
    `]
})
export class LessonComponent {
    @Input() lesson!: Lesson;
    

    constructor(
        private router: Router
    )
    {
    }
    get colorClass(): string {
        console.log('Lesson type:', this.lesson.type); // Для отладки
        
        const colors = {
            lecture: 'bg-blue-100 border-blue-200 hover:bg-blue-200',
            practice: 'bg-green-100 border-green-200 hover:bg-green-200',
            lab: 'bg-purple-100 border-purple-200 hover:bg-purple-200'
        };
        
        const color = colors[this.lesson.type as keyof typeof colors] || 'bg-gray-100 border-gray-200';
        console.log('Selected color:', color); // Для отладки
        
        return color;
    }

    onLessonClick(): void {
        console.log('Lesson clicked:', this.lesson);
        
        // Здесь можно добавить логику при клике на урок
        this.router.navigate(['/distancelearning/view'], { 
            queryParams: { id: this.lesson.id } 
        });
    }
}
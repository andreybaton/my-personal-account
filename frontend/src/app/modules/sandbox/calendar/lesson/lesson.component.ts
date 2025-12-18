import { ChangeDetectionStrategy, Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
//import { Lesson } from './lesson.types';
import { Router } from '@angular/router';
import { Lesson } from '../../../../shared/lesson.interface';
@Component({
    selector: 'app-lesson',
    standalone: true,
    imports: [CommonModule],
    templateUrl: './lesson.component.html',
    styles: [`
        .lesson-card {
        min-height: 60px;
        }
    `],
	changeDetection: ChangeDetectionStrategy.OnPush
})
export class LessonComponent {
    @Input() lesson!: Lesson;
    

    constructor(
        private router: Router
    )
    {
		console.log(this.lesson)
    }
    get colorClass(): string {
        console.log('Lesson type:', this.lesson.lessonType); // Для отладки
        
        const colors = {
            lecture: 'bg-blue-200 hover:bg-blue-300',
            practice: 'bg-green-200 hover:bg-green-300',
            lab: 'bg-orange-200 hover:bg-orange-300'
        };
        
        const color = colors[this.lesson.lessonType as keyof typeof colors] || 'bg-gray-100 border-gray-200';
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
import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';
import { Lesson } from './lesson-view.types';
import { getDayOfWeekName } from './lesson-view.types';

@Component({
    selector: 'app-lesson-view',
    standalone: true,
    imports: [CommonModule],
    templateUrl: './lesson-view.component.html',
    styleUrl: './lesson-view.component.scss'
})
export class LessonViewComponent implements OnInit {
    // @Input() lesson!: Lesson;
    lessonId: number | null = null;
    lesson!: Lesson|undefined;
    lessons: Lesson[] = [
        {
          id: 1,
          title: 'Математика',
          teacher: 'Иванов А.П.',
          time: '9:00-10:30',
          room: 'А-101',
          type: 'lecture',
          day: 1,
          week: 1,  // Добавляем номер недели
          timeSlot: 1
        },
        {
          id: 2,
          title: 'Программирование',
          teacher: 'Петрова С.И.',
          time: '10:45-12:15',
          room: 'Б-205',
          type: 'lab',
          day: 1,
          week: 1,  // Добавляем номер недели
          timeSlot: 2
        },];

    constructor(
        private route: ActivatedRoute,
        private router: Router
    ) 
    {
    }

    ngOnInit(): void {
        this.route.queryParams.subscribe(params => {
        this.lessonId = params['id'] ? Number(params['id']) : null;
        if(this.lessonId){
            this.lesson = this.loadLesson(this.lessonId);
        }
        });
    }
    private loadLesson(lessonId: number): Lesson|undefined {
        return this.lessons.find(lesson => 
            lesson.id === lessonId)
    }

    // getDayOfWeek(): string {
    //     return getDayOfWeekName(this.lesson.day);
    // }
    // Иконки для разных типов занятий
}
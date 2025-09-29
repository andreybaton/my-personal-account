import { Component, Input, OnInit } from '@angular/core';
import { DistlearningService } from './distlearning.service';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { LessonComponent } from './lesson/lesson.component';
import { Lesson } from './lesson/lesson.types';

@Component({
  selector: 'app-distlearning',
  standalone: true,
  imports: [CommonModule, LessonComponent],
  templateUrl: './distlearning.component.html',
//   styleUrl: './distlearning.css'
  
})
export class DistlearningComponent implements OnInit {
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
        private _distservice: DistlearningService,
        private _router: Router,
    )
    {
    }
    ngOnInit(): void {
        console.log('dist learn load')
    }
  
  // Дни недели
    daysOfWeek = [
        { id: 1, name: 'Понедельник', short: 'Пн' },
        { id: 2, name: 'Вторник', short: 'Вт' },
        { id: 3, name: 'Среда', short: 'Ср' },
        { id: 4, name: 'Четверг', short: 'Чт' },
        { id: 5, name: 'Пятница', short: 'Пт' },
        { id: 6, name: 'Суббота', short: 'Сб' },
        { id: 7, name: 'Воскресенье', short: 'Вс' }
    ];

    // Временные промежутки (пары)
    timeSlots = [
        { id: 1, time: '9:00 - 10:30' },
        { id: 2, time: '10:45 - 12:15' },
        { id: 3, time: '13:00 - 14:30' },
        { id: 4, time: '14:45 - 16:15' },
        { id: 5, time: '16:30 - 18:00' },
        { id: 6, time: '18:15 - 19:45' }
    ];

    weeks = [
        { id: 1, name: 'Первая неделя' },
        { id: 2, name: 'Вторая неделя' },
        { id: 3, name: 'Третья неделя' },
        { id: 4, name: 'Четвертая неделя' },
        { id: 5, name: 'Пятая неделя' }
    ];

  // Получить урок для конкретной ячейки
    getLessonsForDayAndWeek(dayId: number, weekId: number): Lesson[] {
        return this.lessons.filter(lesson => 
            lesson.day === dayId && lesson.week === weekId
            ).sort((a, b) => a.timeSlot - b.timeSlot); // Сортируем по времени
    }
}

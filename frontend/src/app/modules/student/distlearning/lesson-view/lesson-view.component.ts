import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';
//import { Lesson } from './lesson-view.types';
//import { getDayOfWeekName } from './lesson-view.types';
import { DistlearningService } from '../distlearning.service';
import { BehaviorSubject, filter, find, map, Subject } from 'rxjs';
import { Lesson } from '../../../../shared/lesson.interface';
import { LessontypeToRusPipe } from '../../../../shared/lessontype-to-rus-pipe';

@Component({
    selector: 'app-lesson-view',
    standalone: true,
    imports: [CommonModule, LessontypeToRusPipe],
    templateUrl: './lesson-view.component.html',
    styleUrl: './lesson-view.component.scss'
})
export class LessonViewComponent implements OnInit {
    lessonId: number | null = null;
    
	less = new Subject<Lesson>();
	less$ = this.less.asObservable();
    constructor(
        private route: ActivatedRoute,
  		private _distservice: DistlearningService
    ) 
    {
    }

    ngOnInit(): void {
        this.route.queryParams.subscribe(params => {
        this.lessonId = params['id'] ? Number(params['id']) : null;
        if(this.lessonId){
        	this._distservice.getAllLessons()
			.pipe(
				map((lessons: Lesson[])=>lessons.find(les=>les.id==this.lessonId)))
				.subscribe(
				(data: Lesson | undefined) => {
					if (data){
						this.less.next(data);
						
					} else{}
				}
			)
        }
        });
    }
    
	formatTime(time: string| undefined): string {
    if (!time) return '';
    return time.substring(0, 5); 
	}
}
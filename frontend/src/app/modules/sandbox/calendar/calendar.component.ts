import { CommonModule } from "@angular/common";
import { ChangeDetectionStrategy, Component, OnInit, ViewChild } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { NzCalendarModule, NzCalendarMode, NzCalendarComponent } from "ng-zorro-antd/calendar";
import { LessonComponent } from "./lesson/lesson.component";
import { Lesson } from "../../../shared/lesson.interface";
import { DistlearningService } from "../../student/distlearning/distlearning.service";

@Component({
  selector: 'calendar',
  standalone: true,
  imports: [FormsModule,CommonModule, NzCalendarModule, LessonComponent ],
  templateUrl: './calendar.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush
//   styleUrl: './distlearning.css'
  
})
export class CalendarComponent implements OnInit {
	@ViewChild('calendarRef') calendarRef!: NzCalendarComponent;
	now = new Date();
	displayedMonth: number = this.now.getMonth();
	lessons: Lesson[] = [];
	lessonsGroupByDay : { [date: string]: Lesson[] } = {};
	constructor( 
		private _distservice: DistlearningService)
	{}

	ngOnInit(): void {
		this._distservice.getAllLessons()
		.subscribe((data)=>{
					
					this.lessons=data;
					this.lessonsGroupByDay = data.reduce((acc, item)=>{
						const key = item.lessonDate;
						if (!acc[key]){
							acc[key]=[];
						}
						acc[key].push(item);
						return acc;
					}, {} as { [date: string]: Lesson[]})
		}
		);
	}

	prevMonth(): void {
		if (this.calendarRef) {
			
			this.calendarRef.onMonthSelect(this.displayedMonth-=1);
		}
	}

	nextMonth(): void {
		if (this.calendarRef) {
			
			this.calendarRef.onMonthSelect(this.displayedMonth+=1);
		}
	}
  
	// getLessonsForDay(date: any){
	// 	const day = date.getDate();
	// 	const month = date.getMonth()+1;
	// 	const year = date.getFullYear();
	// 	const dateStr = `${day}.${month}.${year}`;
	// 	console.log(dateStr);
	// 	return this.lessons.filter(lesson => lesson.lessonDate === dateStr);
	// }

  	panelChange(change: { date: Date; mode: string }): void {
    	console.log(change.date, change.mode);
  	}

	// lessons = [
	// 	{
	// 		id: 1,
	// 		teacher: "Иванов Иван Иванович",
	// 		classroom: "А-101",
	// 		discipline: "Математический анализ",
	// 		lessonType: "lecture",
	// 		lessonDate: "15.12.2025",
	// 		startTime: "8:00",
	// 		endTime: "9:30",
			
	// 	} as Lesson,
	// 	{
	// 		id: 2,
	// 		teacher: "Петров Петр Петрович",
	// 		classroom: "Б-205",
	// 		discipline: "Программирование на C#",
	// 		lessonDate: "15.12.2025",
	// 		startTime: "9:45",
	// 		endTime: "11:20",
	// 		lessonType: "practice"
	// 	}as Lesson,
	// ] 
	// lessonsGroupByDay = this.lessons.reduce((acc, lesson) => {
	// 	const date = lesson.lessonDate;
		
	// 	// Если ключа с такой датой еще нет, создаем пустой массив
	// 	if (!acc[date]) {
	// 		acc[date] = [];
	// 	}
		
	// 	// Добавляем текущее занятие в массив соответствующей даты
	// 	acc[date].push(lesson);
		
	// 	return acc;
	// 	}, {} as { [date: string]: Lesson[] });
	// lessonsGroupByDay: {[date: string]: Lesson[] } = {
	// 	"15.12.2025": [{
	// 		id: 1,
	// 		teacher: "Иванов Иван Иванович",
	// 		classroom: "А-101",
	// 		discipline: "Математический анализ",
	// 		lessonType: "lecture",
	// 		lessonDate: "15.12.2025",
	// 		startTime: "8:00",
	// 		endTime: "9:30",
			
	// 	} as Lesson,
	// 	{
	// 		id: 2,
	// 		teacher: "Петров Петр Петрович",
	// 		classroom: "Б-205",
	// 		discipline: "Программирование на C#",
	// 		lessonDate: "15.12.2025",
	// 		startTime: "9:45",
	// 		endTime: "11:20",
	// 		lessonType: "practice"
	// 	}as Lesson,
	// ]
	// } 

}
export interface Lesson {
    id: number;
    teacher: string;
    classroom: string;
	//lesson: 'lecture' | 'practice' | 'lab';
	discipline: string;
    lessonType: 'lecture' | 'practice' | 'lab';
    lessonDate: number;
	startTime: string;
    endTime: string;   
    //timeSlot:number;
}
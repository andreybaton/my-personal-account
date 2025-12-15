export interface Lesson {
    id: number;
    teacher: string;
    classroom: string;
	//lesson: 'lecture' | 'practice' | 'lab';
    lessonType: 'lecture' | 'practice' | 'lab';
    lessonDate: number;
	startTime: string;
    endTime: number;   
    //timeSlot:number;
}
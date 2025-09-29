export interface Lesson {
    id: number;
    title: string;
    teacher: string;
    time: string;
    room: string;
    type: 'lecture' | 'practice' | 'lab';
    color?: string;
    day: number;
    week: number;   
    timeSlot:number;
}
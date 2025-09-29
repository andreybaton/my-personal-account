export interface Lesson {
  id: number;
  title: string;
  teacher: string;
  time: string;
  room: string;
  type: 'lecture' | 'practice' | 'lab';
  color?: string;
  day: number;        // День недели (1-7)
  week: number;       // Номер недели (1-5)
  timeSlot: number;   // Номер пары (1-6)
  description?: string; // Описание занятия (опционально)
}

// Вспомогательная функция для получения названия дня недели
export function getDayOfWeekName(dayNumber: number): string {
  const days = [
    'Понедельник',
    'Вторник', 
    'Среда',
    'Четверг',
    'Пятница',
    'Суббота',
    'Воскресенье'
  ];
  return days[dayNumber - 1] || 'Неизвестно';
}
-- Создание таблицы lessons
CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    teacher VARCHAR(255) NOT NULL,
    classroom VARCHAR(50) NOT NULL,
    discipline VARCHAR(255) NOT NULL,
    lesson_date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    type VARCHAR(255) NOT NULL
);

-- Вставка тестовых данных
INSERT INTO lessons (teacher, classroom, discipline, lesson_date, start_time, end_time, type) VALUES
('Иванов И.И.', 'А-101', 'Математический анализ', '2024-01-15', '09:00:00', '10:30:00', 'Лекция'),
('Петров П.П.', 'Б-205', 'Программирование', '2024-01-15', '10:45:00', '12:15:00', 'Практика'),
('Сидорова С.С.', 'В-301', 'Базы данных', '2024-01-16', '13:00:00', '14:30:00', 'Лабораторная'),
('Кузнецов К.К.', 'А-102', 'Физика', '2024-01-16', '14:45:00', '16:15:00', 'Лекция'),
('Васильева В.В.', 'Г-401', 'Английский язык', '2024-01-17', '09:00:00', '10:30:00', 'Семинар');

-- Проверка данных
SELECT * FROM lessons;
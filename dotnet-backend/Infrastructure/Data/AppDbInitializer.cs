using Domain.Entities;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;

namespace Infrastructure.Data
{
    public static class DbInitializer
    {
        public static void Initialize(ApplicationDbContext context)
        {
            // Гарантируем, что база создана
            context.Database.EnsureCreated();

            // Проверяем, есть ли уже данные
            if (context.Lessons.Any())
            {
                return; // База уже заполнена
            }

            // Создаем тестовые данные
            var lessons = new Lesson[]
            {
                new Lesson
                {
                    Teacher = "Иванов Иван Иванович",
                    Classroom = "А-101",
                    Discipline = "Математический анализ",
                    LessonDate = new DateOnly(2024, 1, 15), // Только дата
                    StartTime = new TimeOnly(8,0),     // 09:00:00
                    EndTime = new TimeOnly(9,30),     // 10:30:00
                    LessonType = "lecture"
                },
                new Lesson
                {
                    Teacher = "Петров Петр Петрович",
                    Classroom = "Б-205",
                    Discipline = "Программирование на C#",
                    LessonDate = new DateOnly(2024, 1, 15),
                    StartTime = new TimeOnly (9, 45),   // 10:45:00
                    EndTime = new TimeOnly (9, 15),     // 12:15:00
                    LessonType = "practice"
                },
                new Lesson
                {
                    Teacher = "Сидорова Светлана Сергеевна",
                    Classroom = "В-301",
                    Discipline = "Базы данных",
                    LessonDate = new DateOnly(2024, 1, 16),
                    StartTime = new TimeOnly (9, 0),    // 13:00:00
                    EndTime = new TimeOnly (9, 30),     // 14:30:00
                    LessonType = "lab"
                },
                new Lesson
                {
                    Teacher = "Кузнецов Константин Александрович",
                    Classroom = "А-102",
                    Discipline = "Физика",
                    LessonDate = new DateOnly(2024, 1, 16),
                    StartTime = new TimeOnly (9, 45),   // 14:45:00
                    EndTime = new TimeOnly (9, 15),     // 16:15:00
                    LessonType = "lecture"
                },
                new Lesson
                {
                    Teacher = "Васильева Виктория Олеговна",
                    Classroom = "Г-401",
                    Discipline = "Английский язык",
                    LessonDate = new DateOnly(2024, 1, 17),
                    StartTime = new TimeOnly (9, 0),     // 09:00:00
                    EndTime = new TimeOnly (9, 30),     // 10:30:00
                    LessonType = "practice"
                }
            };

            foreach (var lesson in lessons)
            {
                context.Lessons.Add(lesson);
            }

            context.SaveChanges();

            Console.WriteLine("Database seeded with test data.");
        }
    }
}

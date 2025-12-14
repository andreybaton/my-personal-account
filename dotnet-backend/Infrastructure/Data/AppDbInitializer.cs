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
                    LessonDate = new DateTime(2024, 1, 15), // Только дата
                    StartTime = new TimeSpan(9, 0, 0),     // 09:00:00
                    EndTime = new TimeSpan(10, 30, 0),     // 10:30:00
                    LessonType = "Лекция"
                },
                new Lesson
                {
                    Teacher = "Петров Петр Петрович",
                    Classroom = "Б-205",
                    Discipline = "Программирование на C#",
                    LessonDate = new DateTime(2024, 1, 15),
                    StartTime = new TimeSpan(10, 45, 0),   // 10:45:00
                    EndTime = new TimeSpan(12, 15, 0),     // 12:15:00
                    LessonType = "Практика"
                },
                new Lesson
                {
                    Teacher = "Сидорова Светлана Сергеевна",
                    Classroom = "В-301",
                    Discipline = "Базы данных",
                    LessonDate = new DateTime(2024, 1, 16),
                    StartTime = new TimeSpan(13, 0, 0),    // 13:00:00
                    EndTime = new TimeSpan(14, 30, 0),     // 14:30:00
                    LessonType = "Лабораторная"
                },
                new Lesson
                {
                    Teacher = "Кузнецов Константин Александрович",
                    Classroom = "А-102",
                    Discipline = "Физика",
                    LessonDate = new DateTime(2024, 1, 16),
                    StartTime = new TimeSpan(14, 45, 0),   // 14:45:00
                    EndTime = new TimeSpan(16, 15, 0),     // 16:15:00
                    LessonType = "Лекция"
                },
                new Lesson
                {
                    Teacher = "Васильева Виктория Олеговна",
                    Classroom = "Г-401",
                    Discipline = "Английский язык",
                    LessonDate = new DateTime(2024, 1, 17),
                    StartTime = new TimeSpan(9, 0, 0),     // 09:00:00
                    EndTime = new TimeSpan(10, 30, 0),     // 10:30:00
                    LessonType = "Семинар"
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

using Domain.Entities;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Infrastructure.Data
{
    public class ApplicationDbContext : DbContext
    {
        public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options)
            : base(options)
        {
        }

        public DbSet<Lesson> Lessons { get; set; } = null!;

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            // Конвертер для DateTime (сохраняем только дату)
            //var dateConverter = new ValueConverter<DateTime, DateTime>(
            //    v => v.Date,  // Сохраняем только дату
            //    v => DateTime.SpecifyKind(v, DateTimeKind.Unspecified));

            // Конвертер для TimeSpan (сохраняем как время)
            //var timeConverter = new ValueConverter<TimeSpan, TimeSpan>(
            //    v => v,
            //    v => v);
            // Конвертеры для Npgsql 6.0.3
            var dateOnlyConverter = new ValueConverter<DateOnly, DateTime>(
                dateOnly => dateOnly.ToDateTime(TimeOnly.MinValue),
                dateTime => DateOnly.FromDateTime(dateTime));

            var timeOnlyConverter = new ValueConverter<TimeOnly, TimeSpan>(
                timeOnly => timeOnly.ToTimeSpan(),
                timeSpan => TimeOnly.FromTimeSpan(timeSpan));

            modelBuilder.Entity<Lesson>(entity =>
            {
                entity.ToTable("lessons");

                entity.HasKey(e => e.Id);

                entity.Property(e => e.Id)
                    .HasColumnName("id")
                    .ValueGeneratedOnAdd();

                entity.Property(e => e.Teacher)
                    .HasColumnName("teacher")
                    .HasMaxLength(255)
                    .IsRequired();

                entity.Property(e => e.Classroom)
                    .HasColumnName("classroom")
                    .HasMaxLength(50)
                    .IsRequired();

                entity.Property(e => e.Discipline)
                    .HasColumnName("discipline")
                    .HasMaxLength(255)
                    .IsRequired();


                // DateOnly с конвертером
                entity.Property(e => e.LessonDate)
                    .HasColumnName("lesson_date")
                    .HasConversion(dateOnlyConverter)
                    .HasColumnType("date")
                    .IsRequired();

                // TimeOnly с конвертером
                entity.Property(e => e.StartTime)
                    .HasColumnName("start_time")
                    .HasConversion(timeOnlyConverter)
                    .HasColumnType("time")
                    .IsRequired();

                entity.Property(e => e.EndTime)
                    .HasColumnName("end_time")
                    .HasConversion(timeOnlyConverter)
                    .HasColumnType("time")
                    .IsRequired();

                entity.Property(e => e.LessonType)
                    .HasColumnName("type")
                    .HasMaxLength(255)
                    .IsRequired();

                
            });
        }
    }
}

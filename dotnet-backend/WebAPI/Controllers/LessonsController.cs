using Application.Handlers.Queries;
using Domain.Entities;
using Infrastructure.Data;
using MediatR;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using WebAPI.DTO;

namespace WebAPI.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class LessonsController : ControllerBase
    {
        private readonly ApplicationDbContext _context;
        private readonly IMediator _mediator;

        public LessonsController(ApplicationDbContext context, IMediator mediator)
        {
            _context = context;
            _mediator = mediator;
        }

        // GET: api/lessons
        [HttpGet]
        public async Task<ActionResult<IEnumerable<LessonDto>>> GetLessons()
        {

            try {
                var query = new GetLessonsQuery();

                var lessons = await _mediator.Send(query);
                var result = lessons.Select(lesson => new
                {
                    lesson.Id,
                    lesson.Teacher,
                    lesson.Classroom,
                    lesson.Discipline,
                    LessonDate = lesson.LessonDate.ToString("yyyy-MM-dd"),
                    StartTime = lesson.StartTime.ToString("HH:mm:ss"),
                    EndTime = lesson.EndTime.ToString("HH:mm:ss"),
                    lesson.LessonType
                }).ToList();
                return Ok(result);
            }
            catch (Exception ex) { 
                return StatusCode(500, $"Internal server error: {ex.Message}"); 
            }
            //var lessons = await _context.Lessons.ToListAsync();

            //// Маппинг из Entity в DTO
            //return lessons.Select(lesson => new LessonDto
            //{
            //    Id = lesson.Id,
            //    Teacher = lesson.Teacher,
            //    Classroom = lesson.Classroom,
            //    Discipline = lesson.Discipline,
            //    LessonDate = lesson.LessonDate.ToString("yyyy-MM-dd"),
            //    StartTime = lesson.StartTime.ToString("HH:mm:ss"),
            //    EndTime = lesson.EndTime.ToString("HH:mm:ss"),
            //    LessonType = lesson.LessonType
            //}).ToList();
        }

        // GET: api/lessons/5
        [HttpGet("{id}")]
        public async Task<ActionResult<LessonDto>> GetLesson(int id)
        {
            var lesson = await _context.Lessons.FindAsync(id);

            if (lesson == null)
            {
                return NotFound();
            }

            return new LessonDto
            {
                Id = lesson.Id,
                Teacher = lesson.Teacher,
                Classroom = lesson.Classroom,
                Discipline = lesson.Discipline,
                LessonDate = lesson.LessonDate.ToString("yyyy-MM-dd"),
                StartTime = lesson.StartTime.ToString("HH:mm:ss"),
                EndTime = lesson.EndTime.ToString("HH:mm:ss"),
                LessonType = lesson.LessonType
            };
        }

        // POST: api/lessons
        [HttpPost]
        public async Task<ActionResult<Lesson>> PostLesson(LessonDto lessonDto)
        {
            var lesson = new Lesson
            {
                Teacher = lessonDto.Teacher,
                Classroom = lessonDto.Classroom,
                Discipline = lessonDto.Discipline,
                LessonDate = DateOnly.Parse(lessonDto.LessonDate),
                StartTime = TimeOnly.Parse(lessonDto.StartTime),
                EndTime = TimeOnly.Parse(lessonDto.EndTime),
                LessonType = lessonDto.LessonType
            };

            _context.Lessons.Add(lesson);
            await _context.SaveChangesAsync();

            return CreatedAtAction("GetLesson", new { id = lesson.Id }, lesson);
        }
    }
}

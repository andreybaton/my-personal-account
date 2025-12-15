using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Domain.Entities
{
    public class Lesson
    {
        public int Id { get; set; }
        public string Teacher { get; set; } = string.Empty;
        public string Classroom { get; set; } = string.Empty;
        public string Discipline { get; set; } = string.Empty;
        public DateOnly LessonDate { get; set; } // DateTime вместо DateOnly
        public TimeOnly StartTime { get; set; }  // TimeSpan вместо TimeOnly
        public TimeOnly EndTime { get; set; }    // TimeSpan вместо TimeOnly
        public string LessonType { get; set; } = string.Empty;
    }
}

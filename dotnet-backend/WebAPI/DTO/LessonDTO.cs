namespace WebAPI.DTO
{
    public class LessonDto
    {
        public int Id { get; set; }
        public string Teacher { get; set; } = string.Empty;
        public string Classroom { get; set; } = string.Empty;
        public string Discipline { get; set; } = string.Empty;
        public string LessonDate { get; set; } = string.Empty; // string вместо DateOnly
        public string StartTime { get; set; } = string.Empty;  // string вместо TimeOnly
        public string EndTime { get; set; } = string.Empty;    // string вместо TimeOnly
        public string LessonType { get; set; } = string.Empty;
    }
}

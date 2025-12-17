using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Application.Models;
using AutoMapper;
using Domain.Entities;
//using WebAPI.DTO;

namespace Application.Mappings;

public class MappingProfile : Profile
{
    public MappingProfile() {
        CreateMap<Lesson, LessonDto>().ForMember(x => x.Id, opt => opt.MapFrom(x=>x.Id));
        CreateMap<Lesson, LessonDto>().ForMember(x => x.Teacher, opt => opt.MapFrom(x => x.Teacher));
        CreateMap<Lesson, LessonDto>().ForMember(x => x.Classroom, opt => opt.MapFrom(x => x.Classroom));
        CreateMap<Lesson, LessonDto>().ForMember(x => x.Discipline, opt => opt.MapFrom(x => x.Discipline));
        CreateMap<Lesson, LessonDto>().ForMember(x => x.LessonDate, opt => opt.MapFrom(x => x.LessonDate.ToString("yyyy-MM-dd")));
        CreateMap<Lesson, LessonDto>().ForMember(x => x.StartTime, opt => opt.MapFrom(x => x.StartTime.ToString("HH:mm:ss")));
        CreateMap<Lesson, LessonDto>().ForMember(x => x.EndTime, opt => opt.MapFrom(x => x.EndTime.ToString("HH:mm:ss")));
        CreateMap<Lesson, LessonDto>().ForMember(x => x.LessonType, opt => opt.MapFrom(x => x.LessonType));

    }
}

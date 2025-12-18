using Application.Interfaces;
using Application.Models;
using AutoMapper;
using Domain.Entities;
using MediatR;

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;


namespace Application.Handlers.Queries;

public class GetLessonsQuery : IRequest<List<LessonDto>>
{

}
  
public class GetLessonsHandler : IRequestHandler<GetLessonsQuery, List<LessonDto>>
{
    private readonly ILessonRepository _lessonRepository;
    private readonly IMapper _mapper;


    public GetLessonsHandler(ILessonRepository lessonRepository, IMapper mapper)
    {
        _lessonRepository = lessonRepository;
        _mapper = mapper;
    }

    public async Task<List<LessonDto>> Handle(
        GetLessonsQuery request,
        CancellationToken cancellationToken)
    {
        var lessons = await _lessonRepository.GetAllAsync(cancellationToken);
        return _mapper.Map<List<LessonDto>>(lessons);
    }
}


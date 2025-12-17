using Domain.Entities;
using MediatR;

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Application.Handlers.Queries
{
    public class GetLessonsQuery : IRequest<List<Lesson>>
    {

    }
    public interface ILessonQuery
    {
        Task<IEnumerable<Lesson>> GetLessonsAsync(CancellationToken cancellationToken = default);
    }
    public class GetLessonsHandler : IRequestHandler<GetLessonsQuery, List<Lesson>>
    {
        private readonly ILessonQuery _lessonQuery;


        public GetLessonsHandler(ILessonQuery lessonQuery)
        {
            _lessonQuery = lessonQuery;
           
        }

        public async Task<List<Lesson>> Handle(
            GetLessonsQuery request,
            CancellationToken cancellationToken)
        {
            var lessons = await _lessonQuery.GetLessonsAsync(cancellationToken);
            return lessons.ToList();
        }


    }
}

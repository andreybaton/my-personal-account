using Application.Handlers.Queries;
using Application.Interfaces;
using Domain.Entities;
using Infrastructure.Data;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Infrastructure.Repositories;

public class LessonRepository : ILessonRepository
{
    private readonly ApplicationDbContext _context;

    public LessonRepository(ApplicationDbContext context)
    {
        _context = context;
    }

    public async Task<List<Lesson>> GetAllAsync(CancellationToken cancellationToken = default)
    {
        return await _context.Lessons
            .OrderBy(l => l.LessonDate)
            .ThenBy(l => l.StartTime)
            .ToListAsync(cancellationToken);
    }
}

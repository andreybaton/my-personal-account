using Infrastructure.Data;
using Microsoft.EntityFrameworkCore;

//AppContext.SetSwitch("Npgsql.EnableLegacyTimestampBehavior", true);
var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers();

builder.Services.AddDbContext<ApplicationDbContext>(options =>
    options.UseNpgsql(builder.Configuration.GetConnectionString("DefaultConnection")));
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
builder.Services.AddCors(options =>
{
    options.AddPolicy("AllowAngularApp",
        builder =>
        {
            builder.WithOrigins("http://localhost:4200") // Angular dev server
                   .AllowAnyHeader()
                   .AllowAnyMethod()
                   .AllowCredentials();
        });
});

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}
app.UseCors("AllowAngularApp");

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();
// Автоматическое создание базы и заполнение тестовыми данными
using (var scope = app.Services.CreateScope())
{
    try
    {
        var dbContext = scope.ServiceProvider.GetRequiredService<ApplicationDbContext>();

        // Удаляем базу если она существует
        await dbContext.Database.EnsureDeletedAsync();
        Console.WriteLine("Database deleted.");

        // Создаем новую базу
        await dbContext.Database.EnsureCreatedAsync();
        Console.WriteLine("Database created.");

        // Заполняем тестовыми данными
        DbInitializer.Initialize(dbContext);

        Console.WriteLine("Database seeded with test data.");
    }
    catch (Exception ex)
    {
        Console.WriteLine($"An error occurred while resetting database: {ex.Message}");
    }
}
app.Run();

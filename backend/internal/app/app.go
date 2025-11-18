package app

import (
	"log"
	"time"

	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"

	_ "backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	cfg    *config.Config
	router *gin.Engine
	db     *sqlx.DB
}

func New() *App {
	return &App{}
}

func (a *App) Init() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	a.cfg = cfg

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		return err
	}
	a.db = db

	a.setupRouter()

	return nil
}

func (a *App) setupRouter() {
	if a.cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// инициализация репозиториев
	scheduleRepo := repository.NewScheduleRepository(a.db)
	attendanceRepo := repository.NewAttendanceRepository(a.db)
	groupRepo := repository.NewGroupRepository(a.db)
	studentRepo := repository.NewStudentRepository(a.db)

	// инициализация сервисов
	scheduleService := service.NewScheduleService(scheduleRepo)
	attendanceService := service.NewAttendanceService(attendanceRepo)
	groupService := service.NewGroupService(groupRepo)
	studentService := service.NewStudentService(studentRepo)

	// инициализация обработчиков
	scheduleHandler := handler.NewScheduleHandler(scheduleService)
	attendanceHandler := handler.NewAttendanceHandler(attendanceService)
	groupHandler := handler.NewGroupHandler(groupService)
	studentHandler := handler.NewStudentHandler(studentService)

	api := router.Group("/api/v1")
	{
		// Расписание
		schedule := api.Group("/schedule")
		{
			schedule.GET("/students/:id", scheduleHandler.GetStudentSchedule)
			schedule.GET("/students/:id/weekly", scheduleHandler.GetWeeklySchedule)
			schedule.GET("/students/:id/daily", scheduleHandler.GetDailySchedule)
			schedule.GET("/groups/:id", scheduleHandler.GetGroupSchedule)
			schedule.GET("/teachers/:id", scheduleHandler.GetTeacherSchedule)
			schedule.GET("/classrooms/:id", scheduleHandler.GetClassroomSchedule)

			schedule.POST("/lessons", scheduleHandler.CreateLesson)
			schedule.PUT("/lessons/:id", scheduleHandler.UpdateLesson)
			schedule.DELETE("/lessons/:id", scheduleHandler.DeleteLesson)
			schedule.GET("/lessons/:id", scheduleHandler.GetLessonDetails)
		}

		// Посещаемость
		attendance := api.Group("/attendance")
		{
			attendance.POST("/mark", attendanceHandler.MarkAttendance)
			attendance.POST("/mark-batch", attendanceHandler.MarkAttendanceBatch)
			attendance.GET("/students/:id", attendanceHandler.GetStudentAttendance)
			attendance.GET("/students/:id/stats", attendanceHandler.GetAttendanceStats)
			attendance.GET("/lessons/:id", attendanceHandler.GetLessonAttendance)
			attendance.GET("/groups/:id/summary", attendanceHandler.GetGroupAttendanceSummary)
		}

		// Группы
		groups := api.Group("/groups")
		{
			groups.GET("", groupHandler.GetAllGroups)
			groups.GET("/:id", groupHandler.GetGroupDetails)
			groups.GET("/faculty/:faculty_id", groupHandler.GetGroupsByFaculty)
		}

		// Студенты
		students := api.Group("/students")
		{
			students.GET("/:id/profile", studentHandler.GetStudentProfile)
			students.GET("/:id/academic", studentHandler.GetStudentAcademicInfo)
		}

		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":    "ok",
				"timestamp": time.Now().Format(time.RFC3339),
				"service":   "backend",
			})
		})
	}

	a.router = router
}

func (a *App) Run() error {
	log.Printf("Starting server on port %s", a.cfg.ServerPort)
	return a.router.Run(":" + a.cfg.ServerPort)
}

func (a *App) Close() {
	if a.db != nil {
		a.db.Close()
	}
}

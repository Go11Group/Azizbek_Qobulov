package handler

import (
	"ExamNT/storage/postgres"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB         *sql.DB
	Course     *postgres.CourseRepo
	User       *postgres.UsersRepo
	Lesson     *postgres.LessonsRepo
	Enrollment *postgres.EnrollmentRepo
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		DB:         db,
		Course:     postgres.NewCourseRepo(db),
		User:       postgres.NewUsersRepo(db),
		Lesson:     postgres.NewLessonsRepo(db),
		Enrollment: postgres.NewEnrollmentRepo(db),
	}
}

func (h *Handler) SetupRoutes() {
	r := gin.Default()

	r.POST("/user/create", h.CreateUser)
	r.GET("/user", h.GetAllUsers)
	r.GET("/user/get/:id", h.GetUser)
	r.PUT("/user/update", h.UpdateUser)
	r.DELETE("/user/delete/:id", h.DeleteUser)

	r.POST("/course/create", h.CreateCourse)
	r.GET("/course", h.GetAllCourses)
	r.GET("/course/get/:id", h.GetCourse)
	r.PUT("/course/update", h.UpdateCourse)
	r.DELETE("/course/delete/:id", h.DeleteCourse)

	r.POST("/lessons/create", h.CreateLesson)
	r.GET("/lessons", h.GetAllLessons)
	r.GET("/lessons/:id", h.GetLessonByID)
	r.PUT("/lessons/update", h.UpdateLesson)
	r.DELETE("/lessons/:id", h.DeleteLesson)

	r.POST("/enrollments/enroll", h.EnrollUser)
	r.GET("/enrollments", h.GetAllEnrollments)
	r.GET("/enrollments/:id", h.GetEnrollmentByID)
	r.DELETE("/enrollments/:id", h.DeleteEnrollment)

	r.Run(":7070")
}

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

	r.POST("/users/create", h.CreateUser)
	r.GET("/users", h.GetAllUsers)
	r.GET("/users/:id", h.GetUser)
	r.PUT("/users/update", h.UpdateUser)
	r.DELETE("/users/delete/:id", h.DeleteUser)

	r.POST("/course/create", h.CreateCourse)
	r.GET("/course", h.GetAllCourses)
	r.GET("/course/:id", h.GetCourse)
	r.PUT("/course/update", h.UpdateCourse)
	r.DELETE("/course/delete/:id", h.DeleteCourse)

	r.POST("/lessons/create", h.CreateLesson)
	r.GET("/lessons", h.GetAllLessons)
	r.GET("/lessons/:id", h.GetLessonByID)
	r.PUT("/lessons/update", h.UpdateLesson)
	r.DELETE("/lessons/delete/:id", h.DeleteLesson)

	r.POST("/enrollments/enroll", h.EnrollUser)
	r.GET("/enrollments", h.GetAllEnrollments)
	r.GET("/enrollments/:id", h.GetEnrollmentByID)
	r.PUT("/enrollments/update", h.UpdateEnrollment)
	r.DELETE("/enrollments/delete/:id", h.DeleteEnrollment)

	r.GET("/users/:id/courses", h.GetCoursesByUser)
	r.GET("/users/search", h.SearchUsers)
	r.GET("/courses/:id/lessons", h.GetLessonsByCourse)
	r.GET("/courses/:id/enrollments", h.GetEnrolledUsersByCourse)
	r.GET("/courses/popular", h.GetPopularCourses)

	r.Run(":8080")
}

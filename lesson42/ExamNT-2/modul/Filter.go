package modul

type FilterCourse struct {
	Title  string `form:"title"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type FilterUser struct {
	Name     string `form:"name"`
	Birthday string `form:"birthday"`
	Limit    int    `form:"limit"`
	Offset   int    `form:"offset"`
}

type FilterLesson struct {
	LessonID string `form:"lesson_id"`
	CourseID string `form:"course_id"`
	Title    string `form:"title"`
	Limit    int    `form:"limit"`
	Offset   int    `form:"offset"`
}

type FilterEnrollment struct {
	UserID   string `form:"user_id"`
	CourseID string `form:"course_id"`
	Limit    int    `form:"limit"`
	Offset   int    `form:"offset"`
}

type UserFilter struct {
	Name    string `form:"name" json:"name"`
	Email   string `form:"email" json:"email"`
	AgeFrom int    `form:"agefrom" json:"age_from"`
	AgeTo   int    `form:"ageto" json:"age_to"`
}

type PopularCourse struct {
	CourseID         string `json:"course_id"`
	CourseTitle      string `json:"course_title"`
	EnrollmentsCount int    `json:"enrollments_count"`
}

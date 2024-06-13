package modul

type FilterCourse struct {
	CourseID string `json:"course_id"`
	Title    string `json:"title"`
	Limit    int
	Offset   int
}

type FilterUser struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Limit    int
	Offset   int
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
package postgres

import (
	"ExamNT/modul"
	"database/sql"
	"strconv"
	"time"
)

type LessonsRepo struct {
	db *sql.DB
}

func NewLessonsRepo(db *sql.DB) *LessonsRepo {
	return &LessonsRepo{db: db}
}
func (r *LessonsRepo) Create(lesson *modul.Lesson) error {
	query := "INSERT INTO lessons (course_id, title, content, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING lesson_id"
	return r.db.QueryRow(query, lesson.CourseID, lesson.Title, lesson.Content, time.Now(), time.Now(), 0).Scan(&lesson.LessonID)
}

func (r *LessonsRepo) GetByID(id string) (*modul.Lesson, error) {
	lesson := &modul.Lesson{}
	query := "SELECT lesson_id, course_id, title, content, created_at, updated_at, deleted_at FROM lessons WHERE lesson_id = $1"
	err := r.db.QueryRow(query, id).Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
	return lesson, err
}

func (r *LessonsRepo) Update(lesson *modul.Lesson) error {
	query := "UPDATE lessons SET course_id = $1, title = $2, content = $3, updated_at = $4 WHERE lesson_id = $5"
	_, err := r.db.Exec(query, lesson.CourseID, lesson.Title, lesson.Content, time.Now(), lesson.LessonID)
	return err
}

func (r *LessonsRepo) GetAll(f *modul.FilterLesson) ([]modul.Lesson, error) {
	var lessons []modul.Lesson
	query := "SELECT lesson_id, course_id, title, content, created_at, updated_at, deleted_at FROM lessons "
	params := make(map[string]interface{})
	var args []interface{}
	filter := "WHERE true"

	if f.Title != "" {
		params["title"] = f.Title
		filter += " AND title = :title"
	}

	if f.CourseID != "" {
		params["course_id"] = f.CourseID
		filter += " AND course_id = :course_id"
	}

	if f.LessonID != "" {
		params["lesson_id"] = f.LessonID
		filter += " AND lesson_id = :lesson_id"
	}

	limit := ""
	if f.Limit > 0 {
		limit = " LIMIT " + strconv.Itoa(f.Limit)
	}

	offset := ""
	if f.Offset > 0 {
		offset = " OFFSET " + strconv.Itoa(f.Offset)
	}

	query += filter + limit + offset
	query, args = ReplaceQueryParams(query, params)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var lesson modul.Lesson
		if err := rows.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt); err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}

func (r *LessonsRepo) Delete(id string) error {
	query := "UPDATE lessons SET deleted_at = $1 WHERE lesson_id = $2"
	_, err := r.db.Exec(query, time.Now().Unix(), id)
	return err
}

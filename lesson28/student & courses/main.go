<<<<<<< HEAD
package main

import (
	"fmt"
	// "github.com/Go11Group/at_lesson/lesson28/model"
	"github.com/Go11Group/at_lesson/lesson28/storage/postgres"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//Students
	st := postgres.NewStudentRepo(db)
	students, err := st.GetAllStudents()
	if err != nil {
		panic(err)
	}

	student, _ := st.GetByID(students[2].ID)
	for x := range students {
		fmt.Println(students[x])
	}

	fmt.Println(student, "\n")
	//Courses
	cr := postgres.NewCourseRepo(db)
	kurslar, err := cr.GetAllCourses()
	if err != nil {
		panic(err)
	}

	kurs, _ := cr.GetByID(kurslar[2].ID)
	for x := range kurslar {
		fmt.Println(kurslar[x])
	}

	fmt.Println(kurs)
}
=======
package main

import (
	"fmt"
	// "github.com/Go11Group/at_lesson/lesson28/model"
	"github.com/Go11Group/at_lesson/lesson28/storage/postgres"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//Students
	st := postgres.NewStudentRepo(db)
	students, err := st.GetAllStudents()
	if err != nil {
		panic(err)
	}

	student, _ := st.GetByID(students[2].ID)
	for x := range students {
		fmt.Println(students[x])
	}

	fmt.Println(student, "\n")
	//Courses
	cr := postgres.NewCourseRepo(db)
	kurslar, err := cr.GetAllCourses()
	if err != nil {
		panic(err)
	}

	kurs, _ := cr.GetByID(kurslar[2].ID)
	for x := range kurslar {
		fmt.Println(kurslar[x])
	}

	fmt.Println(kurs)
}
>>>>>>> origin/main

<<<<<<< HEAD
package postgres

import (
	"database/sql"
	"example/modul"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type PersonRepo struct {
	db *sql.DB
}

func NewPersonRepo(db *sql.DB) *PersonRepo {
	return &PersonRepo{db: db}
}

func (r *PersonRepo) GetAll(c *gin.Context) {
	var p modul.Person
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := make(map[string]interface{})
	var args []interface{}
	filter := "WHERE true"

	if len(p.Gender) > 0 {
		params["gender"] = p.Gender
		filter += " AND gender = :gender"
	}

	if len(p.Nation) > 0 {
		params["nation"] = p.Nation
		filter += " AND nation = :nation"
	}

	if len(p.Field) > 0 {
		params["field"] = p.Field
		filter += " AND field = :field"
	}

	if p.Age > 0 {
		params["age"] = p.Age
		filter += " AND age = :age"
	}

	query := `SELECT id, first_name, last_name, age, gender, nation, field, parent_name, city 
              FROM persons ` + filter

	fmt.Println(query)
	query, args = ReplaceQueryParams(query, params)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(query)

	defer rows.Close()

	var persons []modul.Person
	for rows.Next() {
		var person modul.Person
		if err := rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Age, &person.Gender, &person.Nation, &person.Field, &person.ParentName, &person.City); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		persons = append(persons, person)
	}

	c.JSON(http.StatusOK, persons)
}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var args []interface{}
	i := 1

	for k, v := range params {
		holder := ":" + k
		if strings.Contains(namedQuery, holder) {
			namedQuery = strings.ReplaceAll(namedQuery, holder, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}
=======
package postgres

import (
	"database/sql"
	"example/modul"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type PersonRepo struct {
	db *sql.DB
}

func NewPersonRepo(db *sql.DB) *PersonRepo {
	return &PersonRepo{db: db}
}

func (r *PersonRepo) GetAll(c *gin.Context) {
	var p modul.Person
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := make(map[string]interface{})
	var args []interface{}
	filter := "WHERE true"

	if len(p.Gender) > 0 {
		params["gender"] = p.Gender
		filter += " AND gender = :gender"
	}

	if len(p.Nation) > 0 {
		params["nation"] = p.Nation
		filter += " AND nation = :nation"
	}

	if len(p.Field) > 0 {
		params["field"] = p.Field
		filter += " AND field = :field"
	}

	if p.Age > 0 {
		params["age"] = p.Age
		filter += " AND age = :age"
	}

	query := `SELECT id, first_name, last_name, age, gender, nation, field, parent_name, city 
              FROM persons ` + filter

	fmt.Println(query)
	query, args = ReplaceQueryParams(query, params)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(query)

	defer rows.Close()

	var persons []modul.Person
	for rows.Next() {
		var person modul.Person
		if err := rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Age, &person.Gender, &person.Nation, &person.Field, &person.ParentName, &person.City); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		persons = append(persons, person)
	}

	c.JSON(http.StatusOK, persons)
}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var args []interface{}
	i := 1

	for k, v := range params {
		holder := ":" + k
		if strings.Contains(namedQuery, holder) {
			namedQuery = strings.ReplaceAll(namedQuery, holder, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}
>>>>>>> origin/main

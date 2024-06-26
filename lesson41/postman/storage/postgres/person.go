package postgres

import (
	"database/sql"
	"person/model"
	"github.com/pkg/errors"
)

type PersonRepo struct {
	DB *sql.DB
}

func NewPersonRepo(db *sql.DB) *PersonRepo {
	return &PersonRepo{DB: db}
}

func (p *PersonRepo) Create(person model.Person) error {
	if person.Name == "" || person.Age < 1 {
		return errors.New("Empty fields")
	}
	_, err := p.DB.Exec("insert into person(name, age) values($1, $2)",
		person.Name, person.Age)
	if err != nil {
		return errors.Wrap(err, "failed to insert person")
	}
	return nil
}

func (p *PersonRepo) Read(id int) (*model.Person, error) {
	var person model.Person
	row := p.DB.QueryRow("select id, name, age from person where id = $1", id)
	err := row.Scan(&person.ID, &person.Name, &person.Age)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read person")
	}
	return &person, nil
}

func (p *PersonRepo) Update(id int, person model.Person) error {
	query := "update person set name = $1, age = $2 where id = $3"
	_, err := p.DB.Exec(query, person.Name, person.Age, id)
	if err != nil {
		return errors.Wrap(err, "failed to update person")
	}
	return nil
}

func (p *PersonRepo) Delete(id int) error {
	_, err := p.DB.Exec("DELETE from person where id = $1", id)
	if err != nil {
		return errors.Wrap(err, "Failed to delete person")
	}
	return nil
}

func (p *PersonRepo) ReadAll() ([]model.Person, error) {
	rows, err := p.DB.Query("select id, name, age from person")
	if err != nil {
		rows.Close()
		return nil, errors.Wrap(err, "failed to retrieve people")
	}
	defer rows.Close()

	var people []model.Person
	for rows.Next() {
		var person model.Person
		err = rows.Scan(&person.ID, &person.Name, &person.Age)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read")
		}
		people = append(people, person)
	}
	return people, nil
}

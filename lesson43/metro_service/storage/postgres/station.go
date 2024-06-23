package postgres

import (
	"database/sql"
	"github.com/Go11Group/at_lesson/lesson43/metro_service/models"
	"github.com/google/uuid"
)

type StationRepo struct {
	Db *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{Db: db}
}

func (s *StationRepo) Create(station *models.Station) error {
	_, err := s.Db.Exec("insert into station(id, name) values ($1, $2)", uuid.NewString(), station.Name)
	return err
}

func (s *StationRepo) GetById(id string) (*models.Station, error) {
	var station = models.Station{Id: id}
	err := s.Db.QueryRow("select name from station where id = $1", id).Scan(&station.Name)
	if err != nil {
		return nil, err
	}
	return &station, nil
}



func (c *StationRepo) Update(id string, station *models.Station) error {
	_, err := c.Db.Exec("UPDATE station SET number = $1, user_id = $2, level_id = $3 WHERE id = $4", station.Name, id)
	return err
}

func (c *StationRepo) Delete(id string) error {
	_, err := c.Db.Exec("DELETE FROM station WHERE id = $1", id)
	return err
}
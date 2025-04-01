package storage

import (
	"fmt"
	"strconv"

	"github.com/fevse/effm/internal/config"
	"github.com/fevse/effm/internal/logger"
	_ "github.com/jackc/pgx/stdlib" // driver
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

type Storage struct {
	db     *sqlx.DB
	config *config.Config
	logger *logger.Logger
}

func NewStorage(config *config.Config, logger *logger.Logger) *Storage {
	return &Storage{
		config: config,
		logger: logger,
	}
}

func (s *Storage) Migrate() error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("migration set dialect error: %w", err)
	}
	if err := goose.Up(s.db.DB, "."); err != nil {
		return fmt.Errorf("migration up error: %w", err)
	}
	s.logger.Info("migration completed successfully")
	return nil
}

func (s *Storage) Connect() (err error) {
	dsn := s.config.DBConnectionString()

	s.db, err = sqlx.Connect("pgx", dsn)
	if err != nil {
		return fmt.Errorf("connection db error: %w", err)
	}
	s.logger.Info("connection completed successfully")
	return nil
}

func (s *Storage) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return fmt.Errorf("db is not open")
}

func (s *Storage) Create(person *Person) error {
	query := `
	INSERT INTO people (name, surname, patronymic, age, sex, nationality)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`
	_, err := s.db.Exec(query, person.Name, person.Surname, person.Patronymic, person.Age, person.Sex, person.Nationality)
	if err != nil {
		s.logger.Error("error creating Person")
		return err
	}
	s.logger.Info("new Person created successfully")
	return nil
}

func (s *Storage) Show(filter map[string]string, limit, offset int) ([]Person, error) {

	query := `
	SELECT id, name, surname, 
	COALESCE(patronymic, '') as patronymic, 
	COALESCE(age, 0) as age, 
	COALESCE(sex, '') as sex, 
	COALESCE(nationality, '') as nationality 
	FROM people
	WHERE 1=1`

	args := []any{}
	counter := 1

	for k, v := range filter {
		query += " AND " + k + " = $" + strconv.Itoa(counter)
		args = append(args, v)
		counter++
	}
	if limit >= 0 {
		query += " LIMIT $" + strconv.Itoa(counter)
		counter++
		args = append(args, limit)
	}
	query += " OFFSET $" + strconv.Itoa(counter)
	args = append(args, offset)

	data := make([]Person, 0)

	err := s.db.Select(&data, query, args...)
	if err != nil {
		return nil, fmt.Errorf("cannot select people from db: %w", err)
	}
	s.logger.Info("People shown successfully")
	return data, nil
}

func (s *Storage) Delete(id int) error {

	query := `DELETE FROM people WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("cannot delete person with id #%d: %w", id, err)
	}
	s.logger.Info("Person " + strconv.Itoa(id) + " deleted successfully")
	return nil
}

func (s *Storage) Update(id int, person *Person) error {

	if person.Name != "" {
		query := `
		UPDATE people
		SET name = $1
		WHERE id = $2`

		_, err := s.db.Exec(query, person.Name, id)
		if err != nil {
			return fmt.Errorf("cannot update person with id #%d: %w", id, err)
		}
	}
	if person.Surname != "" {
		query := `
		UPDATE people
		SET surname = $1
		WHERE id = $2`

		_, err := s.db.Exec(query, person.Surname, id)
		if err != nil {
			return fmt.Errorf("cannot update person with id #%d: %w", id, err)
		}
	}
	if person.Patronymic != "" {
		query := `
		UPDATE people
		SET patronymic = $1
		WHERE id = $2`

		_, err := s.db.Exec(query, person.Patronymic, id)
		if err != nil {
			return fmt.Errorf("cannot update person with id #%d: %w", id, err)
		}
	}
	if person.Age != 0 {
		query := `
		UPDATE people
		SET age = $1
		WHERE id = $2`

		_, err := s.db.Exec(query, person.Age, id)
		if err != nil {
			return fmt.Errorf("cannot update person with id #%d: %w", id, err)
		}
	}
	if person.Sex != "" {
		query := `
		UPDATE people
		SET sex = $1
		WHERE id = $2`

		_, err := s.db.Exec(query, person.Sex, id)
		if err != nil {
			return fmt.Errorf("cannot update person with id #%d: %w", id, err)
		}
	}
	if person.Nationality != "" {
		query := `
		UPDATE people
		SET nationality = $1
		WHERE id = $2`

		_, err := s.db.Exec(query, person.Nationality, id)
		if err != nil {
			return fmt.Errorf("cannot update person with id #%d: %w", id, err)
		}
	}
	s.logger.Info("Person " + strconv.Itoa(id) + " updated successfully")
	return nil
}

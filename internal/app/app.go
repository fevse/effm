package app

import (
	"encoding/json"
	"net/http"

	"github.com/fevse/effm/internal/config"
	"github.com/fevse/effm/internal/storage"
)

type EffmApp struct {
	Storage Storage
	Logger  Logger
	Config  *config.Config
}

type Storage interface {
	Show(map[string]string, int, int) ([]storage.Person, error)
	Create(*storage.Person) error
	Delete(int) error
	Update(int, *storage.Person) error
}

type Logger interface {
	Debug(string)
	Info(string)
	Error(string)
}

func NewEffmApp(config *config.Config, storage Storage, logger Logger) *EffmApp {
	return &EffmApp{
		Storage: storage,
		Logger:  logger,
		Config:  config,
	}
}

func (e *EffmApp) Show(filter map[string]string, limit, offset int) ([]storage.Person, error) {
	return e.Storage.Show(filter, limit, offset)
}

func (e *EffmApp) Create(person *storage.Person) error {
	age := storage.Age{}
	sex := storage.Sex{}
	nationality := storage.Nationality{}
	name := person.Name

	resp, err := http.Get(e.Config.APIAge + name)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(resp.Body).Decode(&age); err != nil {
		return err
	}

	resp, err = http.Get(e.Config.APISex + name)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(resp.Body).Decode(&sex); err != nil {
		return err
	}

	resp, err = http.Get(e.Config.APINationality + name)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(resp.Body).Decode(&nationality); err != nil {
		return err
	}

	person.Age = age.Age
	person.Sex = sex.Sex
	person.Nationality = nationality.Nationality[0].Country

	return e.Storage.Create(person)
}

func (e *EffmApp) Delete(id int) error {
	return e.Storage.Delete(id)
}

func (e *EffmApp) Update(id int, person *storage.Person) error {
	return e.Storage.Update(id, person)
}

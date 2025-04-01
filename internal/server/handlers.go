package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fevse/effm/internal/storage"
)

// Show godoc
// @Summary Get filtered and paginated list of persons
// @Description Retrieves a list of persons with optional filtering and pagination
// @Tags person
// @Produce json
// @Param name query string false "Filter by name"
// @Param limit query int false "Number of items per page (default: all)"
// @Param offset query int false "Offset for pagination (default: 0)"
// @Success 200 {array} []storage.Person
// @Failure 500 {object} nil "Failed to get data"
// @Router / [get]
func (s *Server) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		filter := make(map[string]string)
		var limit, offset int
		for k, v := range r.URL.Query() {
			if k != "limit" && k != "offset" {
				filter[k] = v[0]
			}
		}
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			s.app.Logger.Debug(err.Error())
			limit = -1
		}
		offset, err = strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			s.app.Logger.Debug(err.Error())
			offset = 0
		}

		list, err := s.app.Storage.Show(filter, limit, offset)
		if err != nil {
			s.app.Logger.Error(err.Error())
			http.Error(w, "Failed to get data", http.StatusInternalServerError)
			return
		}

		s.app.Logger.Debug("handler Show, method " + r.Method)
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(list)
		if err != nil {
			s.app.Logger.Error(err.Error())
		}
	}
}

// Create godoc
// @Summary Create a new person
// @Description Creates a new person record
// @Tags person
// @Accept json
// @Produce json
// @Param person body storage.Person true "Person data to create"
// @Success 201 {object} storage.Person
// @Failure 400 {object} nil "Invalid JSON"
// @Failure 500 {object} nil "Failed to create person"
// @Router / [post]
func (s *Server) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var person storage.Person
		if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
			s.app.Logger.Error(err.Error())
			http.Error(w, "Invalid JSON", http.StatusInternalServerError)
			return
		}

		if err := s.app.Create(&person); err != nil {
			s.app.Logger.Error(err.Error())
			http.Error(w, "Failed to create person", http.StatusInternalServerError)
			return
		}

		s.app.Logger.Debug("handler Create, method " + r.Method)
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(person)
		if err != nil {
			s.app.Logger.Error(err.Error())
		}
	}
}

// Delete godoc
// @Summary Delete a person
// @Description Deletes a person by ID
// @Tags person
// @Param id path int true "Person ID"
// @Success 204
// @Failure 400 {object} nil "Invalid ID"
// @Failure 500 {object} nil "Failed to delete person"
// @Router /{id} [delete]
func (s *Server) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			s.app.Logger.Error(err.Error())
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		if err := s.app.Delete(id); err != nil {
			s.app.Logger.Error(err.Error())
			http.Error(w, "Failed to delete person", http.StatusInternalServerError)
			return
		}

		s.app.Logger.Debug("handler Delete, method " + r.Method)
		w.WriteHeader(http.StatusNoContent)
	}
}

// Update godoc
// @Summary Update a person
// @Description Updates an existing person by ID
// @Tags person
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param person body storage.Person true "Updated person data"
// @Success 200 {object} storage.Person
// @Failure 400 {object} nil "Invalid ID"
// @Failure 500 {object} nil "Invalid JSON"
// @Failure 500 {object} nil "Failed to update person"
// @Router /{id} [put]
func (s *Server) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var person storage.Person
		if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
			s.app.Logger.Error(err.Error())
			http.Error(w, "Invalid JSON", http.StatusInternalServerError)
			return
		}

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			s.app.Logger.Error(err.Error())
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := s.app.Update(id, &person); err != nil {
			s.app.Logger.Error(err.Error())
			http.Error(w, "Failed to update person", http.StatusInternalServerError)
			return
		}

		s.app.Logger.Debug("handler Update, method " + r.Method)
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(person)
		if err != nil {
			s.app.Logger.Error(err.Error())
		}
	}
}

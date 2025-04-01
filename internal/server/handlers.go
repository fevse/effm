package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fevse/effm/internal/storage"
)

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
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(person)
		if err != nil {
			s.app.Logger.Error(err.Error())
		}
	}
}

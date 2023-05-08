package handlers

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vsPEach/Kursovaya/internal/entity"
)

type Logger interface {
	Error(...interface{})
	Info(...interface{})
}

type Storage interface {
	Create(ctx context.Context, event entity.Event) error
	Update(ctx context.Context, id uuid.UUID, event entity.Event) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetOnDay(ctx context.Context, date time.Time) ([]entity.Event, error)
	GetOnWeek(ctx context.Context, date time.Time) ([]entity.Event, error)
	GetOnMonth(ctx context.Context, date time.Time) ([]entity.Event, error)
}

func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	event, err := extractEvent(r)
	if err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.storage.Create(r.Context(), event); err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(event); err != nil {
		h.logger.Error(err)
	}
}

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
	event, err := extractEvent(r)
	if err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.storage.Update(r.Context(), event.ID, event); err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(event); err != nil {
		h.logger.Error(err)
	}
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	UUID, err := uuid.Parse(id)
	if err != nil {
		h.logger.Error(err)
	}

	if err := h.storage.Delete(r.Context(), UUID); err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) GetOnDay(w http.ResponseWriter, r *http.Request) {
	events, err := h.storage.GetOnDay(r.Context(), extractTime(r.FormValue("time")))
	if err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.logger.Info(events)
	if err := json.NewEncoder(w).Encode(events); err != nil {
		h.logger.Error(err)
	}
}

func (h *Handlers) GetOnWeek(w http.ResponseWriter, r *http.Request) {
	events, err := h.storage.GetOnWeek(r.Context(), extractTime(r.FormValue("time")))
	if err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(events); err != nil {
		h.logger.Error(err)
	}
}

func (h *Handlers) GetOnMonth(w http.ResponseWriter, r *http.Request) {
	events, err := h.storage.GetOnMonth(r.Context(), extractTime(r.FormValue("time")))
	if err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(events); err != nil {
		h.logger.Error(err)
	}
}

type Handlers struct {
	logger  Logger
	storage Storage
}

func NewHTTPHandlers(logger Logger, storage Storage) *Handlers {
	return &Handlers{logger: logger, storage: storage}
}

func (h *Handlers) Routes() *mux.Router {
	r := mux.NewRouter().PathPrefix("/calendar").Subrouter()
	r.HandleFunc("/update", h.Update).Methods(http.MethodPost)
	r.HandleFunc("/create", h.Create).Methods(http.MethodPost)
	r.HandleFunc("/delete/{id}", h.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/read/day", h.GetOnDay).Methods(http.MethodGet).Queries("time", "{\\d{4}(.\\d{2}){2}(\\ s|T)(\\d{2}.){2}\\d{2}}")
	r.HandleFunc("/read/week", h.GetOnWeek).Methods(http.MethodGet).Queries("time", "{\\d{4}(.\\d{2}){2}(\\ s|T)(\\d{2}.){2}\\d{2}}")
	r.HandleFunc("/read/month", h.GetOnMonth).Methods(http.MethodGet).Queries("time", "{\\d{4}(.\\d{2}){2}(\\ s|T)(\\d{2}.){2}\\d{2}}")
	return r
}

func extractEvent(r *http.Request) (entity.Event, error) {
	var event entity.Event
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return event, err
	}

	err = json.Unmarshal(data, &event)
	if err != nil {
		return event, err
	}

	return event, nil
}

func extractTime(Time string) time.Time {
	parse, err := time.Parse("02-01-2006T15:04:05", Time)
	if err != nil {
		return time.Time{}
	}
	return parse
}

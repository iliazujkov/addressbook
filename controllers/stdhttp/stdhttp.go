package stdhttp

import (
	"CryptoLAB/gate/psg"
	"CryptoLAB/models/dto"
	"CryptoLAB/pkg"
	"encoding/json"
	"net/http"
)

// Controller обрабатывает HTTP запросы для адресной книги.
type Controller struct {
	DB  *psg.Psg
	Srv *http.Server
}

// NewController создает новый Controller.
func NewController(addr string, db *psg.Psg) *Controller {
	return &Controller{
		DB: db,
		Srv: &http.Server{
			Addr: addr,
		},
	}
}

// RecordAdd обрабатывает HTTP запрос для добавления новой записи.
func (c *Controller) RecordAdd(w http.ResponseWriter, r *http.Request) {
	var record dto.Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	normalizedPhone, err := pkg.PhoneNormalize(record.Phone)
	if err != nil {
		http.Error(w, "Error normalizing phone number", http.StatusInternalServerError)
		return
	}
	record.Phone = normalizedPhone

	id, err := c.DB.RecordAdd(record)
	if err != nil {
		http.Error(w, "Error adding record to the database", http.StatusInternalServerError)
		return
	}

	response := map[string]int64{"id": id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RecordsGet обрабатывает HTTP запрос для получения записей на основе предоставленных полей Record.
func (c *Controller) RecordsGet(w http.ResponseWriter, r *http.Request) {
	var record dto.Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	records, err := c.DB.RecordsGet(record)
	if err != nil {
		http.Error(w, "Error retrieving records from the database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

// RecordUpdate обрабатывает HTTP запрос для обновления записи.
func (c *Controller) RecordUpdate(w http.ResponseWriter, r *http.Request) {
	var record dto.Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	normalizedPhone, err := pkg.PhoneNormalize(record.Phone)
	if err != nil {
		http.Error(w, "Error normalizing phone number", http.StatusInternalServerError)
		return
	}
	record.Phone = normalizedPhone

	err = c.DB.RecordUpdate(record)
	if err != nil {
		http.Error(w, "Error updating record in the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// RecordDeleteByPhone обрабатывает HTTP запрос для удаления записи по номеру телефона.
func (c *Controller) RecordDeleteByPhone(w http.ResponseWriter, r *http.Request) {
	var record dto.Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	normalizedPhone, err := pkg.PhoneNormalize(record.Phone)
	if err != nil {
		http.Error(w, "Error normalizing phone number", http.StatusInternalServerError)
		return
	}
	record.Phone = normalizedPhone

	err = c.DB.RecordDeleteByPhone(record.Phone)
	if err != nil {
		http.Error(w, "Error deleting record from the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

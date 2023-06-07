package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rodrigosollberger/sheikah-slate/internal/app/creature"
)

// CreatureHandler handles HTTP requests for creature operations
type CreatureHandler struct {
	repo creature.Repository
}

// NewCreatureHandler creates a new CreatureHandler instance
func NewCreatureHandler(repo creature.Repository) *CreatureHandler {
	return &CreatureHandler{repo: repo}
}

// GetAllCreatures returns all creatures
func (h *CreatureHandler) GetAllCreatures(w http.ResponseWriter, r *http.Request) {
	creatures, err := h.repo.GetAll()
	if err != nil {
		log.Printf("Failed to get creatures: %v", err)
		http.Error(w, "Failed to get creatures", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(creatures)
}

// GetCreatureByID returns a creature by ID
func (h *CreatureHandler) GetCreatureByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	creatureID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid creature ID", http.StatusBadRequest)
		return
	}

	creature, err := h.repo.GetByID(creatureID)
	if err != nil {
		log.Printf("Failed to get creature with ID %d: %v", creatureID, err)
		http.Error(w, "Failed to get creature", http.StatusInternalServerError)
		return
	}
	if creature == nil {
		http.Error(w, "Creature not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(creature)
}

// CreateCreature creates a new creature
func (h *CreatureHandler) CreateCreature(w http.ResponseWriter, r *http.Request) {
	var creature creature.Creature
	err := json.NewDecoder(r.Body).Decode(&creature)
	if err != nil {
		http.Error(w, "Invalid creature data", http.StatusBadRequest)
		return
	}

	err = h.repo.Create(&creature)
	if err != nil {
		log.Printf("Failed to create creature: %v", err)
		http.Error(w, "Failed to create creature", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(creature)
}

// UpdateCreature updates an existing creature
func (h *CreatureHandler) UpdateCreature(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	creatureID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid creature ID", http.StatusBadRequest)
		return
	}

	var creature creature.Creature
	err = json.NewDecoder(r.Body).Decode(&creature)
	if err != nil {
		http.Error(w, "Invalid creature data", http.StatusBadRequest)
		return
	}

	err = h.repo.Update(creatureID, &creature)
	if err != nil {
		log.Printf("Failed to update creature with ID %d: %v", creatureID, err)
		http.Error(w, "Failed to update creature", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(creature)
}

// DeleteCreature deletes a creature
func (h *CreatureHandler) DeleteCreature(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	creatureID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid creature ID", http.StatusBadRequest)
		return
	}

	err = h.repo.Delete(creatureID)
	if err != nil {
		log.Printf("Failed to delete creature with ID %d: %v", creatureID, err)
		http.Error(w, "Failed to delete creature", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// RegisterRoutes registers the creature routes
func RegisterRoutes(router *mux.Router, repo creature.Repository) {
	handler := NewCreatureHandler(repo)

	router.HandleFunc("/creatures", handler.GetAllCreatures).Methods("GET")
	router.HandleFunc("/creatures/{id}", handler.GetCreatureByID).Methods("GET")
	router.HandleFunc("/creatures", handler.CreateCreature).Methods("POST")
	router.HandleFunc("/creatures/{id}", handler.UpdateCreature).Methods("PUT")
	router.HandleFunc("/creatures/{id}", handler.DeleteCreature).Methods("DELETE")
}

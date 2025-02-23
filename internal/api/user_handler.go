package api

import (
	"encoding/json"
	"log"
	"net/http"

	"simple.market/internal/domain"
	"simple.market/internal/repository"
	"simple.market/internal/service"
	"simple.market/pkg/utils"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var data domain.SignUp

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error to decode JSON", http.StatusBadRequest)
		return
	}

	dbPath := "../../database.db"
	db, err := utils.GetConnection(dbPath)
	if err != nil {
		log.Fatal("Error to connect db", err)
	}
	defer db.Close()

	repo := repository.NewUserRepositorySQLite(db)

	userService := service.NewUserService(repo)

	_, err = userService.CreateUser(data.Email, data.Password, data.PasswordConfirmation)
	if err != nil {
		http.Error(w, "Failed to sign up user", http.StatusInternalServerError)
	}

	// FIX ME
	w.WriteHeader(http.StatusOK)
}

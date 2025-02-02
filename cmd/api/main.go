package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/POST /signup", Signup)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Server listening on port 3000")
	server.ListenAndServe()
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var signUpData SignUp

	err := json.NewDecoder(r.Body).Decode(&signUpData)
	if err != nil {
		http.Error(w, "Error to decode JSON", http.StatusBadRequest)
		return
	}

	isValid := isValidEmail(signUpData.Email)
	if isValid != true {
		http.Error(w, "Invalid email", http.StatusBadRequest)
	}

	if signUpData.Password != signUpData.PasswordConfirmation {
		http.Error(w, "Password and password confirmation doesnt match", http.StatusBadRequest)
	}

	hashedPassword, err := HashedPassword(signUpData.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
	}

	result, err := SignUpService.create(signUpData.Email, hashedPassword)
	if err != nil {
		http.Error(w, "Failed to sign up user", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"uuid": result.uuid, "email": result.email})

}

func HashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(passsword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passsword))
	return err == nil
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUp struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

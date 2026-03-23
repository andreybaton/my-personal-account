package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo  *UserRepository
	jwtSecret []byte
}

func NewAuthHandler(userRepo *UserRepository, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		userRepo:  userRepo,
		jwtSecret: []byte(jwtSecret),
	}
}

func (h *AuthHandler) registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	json.NewDecoder(r.Body).Decode(&req)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user := &User{
		Email:     req.Email,
		Password:  string(hashedPass),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
	}

	if err := h.userRepo.Create(user); err != nil {
		//проверка, мб почта существует
		http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"email":      user.Email,
		"role":       user.Role,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(h.jwtSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	response := AuthResponse{
		Token:     tokenString,
		UserID:    user.ID,
		Email:     user.Email,
		Role:      user.Role,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

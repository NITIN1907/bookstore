package handlers

import (
	"backend_development_go/config"
	"backend_development_go/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)
	config.DB.Create(&user)
	json.NewEncoder(w).Encode(map[string]string{"status": "register"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	json.NewDecoder(r.Body).Decode(&creds)

	var dbUser models.User
	if err := config.DB.Where("email=?", creds.Email).First(&dbUser).Error; err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": dbUser.ID,
		"role":    dbUser.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte("secret"))
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

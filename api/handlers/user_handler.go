package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alifhaider/carehub-backend/internal/database"
	"github.com/alifhaider/carehub-backend/internal/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    database.DB.Find(&users)
    json.NewEncoder(w).Encode(users)
}
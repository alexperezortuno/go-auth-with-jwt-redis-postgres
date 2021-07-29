package auth

import (
	"encoding/json"
	db "github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/data_base"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/data_base/model"
	"log"
)

func UnmarshalAuthRequest(data []byte) (AuthRequest, error) {
	var r AuthRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuthRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ValidateUser(ar AuthRequest) (bool, string) {
	user := model.User{}
	verifyUser, err := user.ValidUser(db.Connection, ar.Email, ar.Password)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return false, err.Error()
	}

	log.Printf("%s", verifyUser)
	return true, ""
}

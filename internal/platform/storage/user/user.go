package user

import (
	"encoding/json"
	db "github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/data_base"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/data_base/model"
	"github.com/jinzhu/copier"
	"log"
	"strings"
)

func UnmarshalInputRequest(data []byte) (InputRequest, error) {
	var r InputRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InputRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type InputRequest struct {
	FullName *string `json:"full_name,omitempty"`
	Name     *string `json:"name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	IDCard   *string `json:"id_card,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
}

func CreateNewUser(ir InputRequest) (model.UserResponse, string) {
	user := model.User{}
	userCopy := model.UserResponse{}
	var errResp string
	var userCreated *model.User
	data, err := ir.Marshal()

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}

	user.Prepare()

	err = user.Validate("")
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		errResp = err.Error()
	}

	userCreated, err = user.SaveUser(db.Connection)

	if err != nil {
		log.Printf("Error %s", err.Error())
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return userCopy, "user already exists"
		}
	}

	err = copier.Copy(&userCopy, userCreated)
	if err != nil {
		return userCopy, "internal server error"
	}

	return userCopy, errResp
}

func GetById(id int) (model.UserResponse, string) {
	user := model.User{}
	userCopy := model.UserResponse{}
	var errResp string

	findUser, err := user.FindUserByID(db.Connection, id)
	if err != nil {
		log.Printf("Error %s", err.Error())
	}

	err = copier.Copy(&userCopy, findUser)
	if err != nil {
		log.Printf("Error %s", err.Error())
		return userCopy, ""
	}

	return userCopy, errResp
}

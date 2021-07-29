package model

import (
	"encoding/json"
	"errors"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"log"
	"strings"
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FullName  string    `gorm:"size:255;null;" json:"full_name"`
	Name      string    `gorm:"size:100;not null;" json:"name"`
	LastName  string    `gorm:"size:100;not null;" json:"last_name"`
	Nickname  string    `gorm:"size:50;null;unique" json:"nickname"`
	IdCard    string    `gorm:"size:30;not null;unique" json:"id_card"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserResponse struct {
	ID       uint32 `json:"id" copier:"must"`
	FullName string `json:"full_name"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Nickname string `json:"nickname"`
	IdCard   string `json:"id_card"`
	Email    string `json:"email" copier:"must"`
}

func UnmarshalUser(data []byte) (User, error) {
	var r User
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *User) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.IdCard = html.EscapeString(strings.TrimSpace(u.IdCard))
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.FullName = html.EscapeString(u.Name + " " + u.LastName)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("required name")
		}
		if u.LastName == "" {
			return errors.New("required last name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("required name")
		}
		if u.LastName == "" {
			return errors.New("required last name")
		}
		if u.IdCard == "" {
			return errors.New("required id card")
		}
		if u.Nickname == "" {
			return errors.New("required nickname")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error

	err = u.BeforeSave()
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}

	err = db.Debug().Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}

	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid int) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error

	if err != nil {
		return &User{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &User{}, errors.New("user not found")
	}

	return u, err
}

func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"nickname":   u.Nickname,
			"email":      u.Email,
			"updated_at": time.Now(),
		},
	)

	if db.Error != nil {
		return &User{}, db.Error
	}

	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

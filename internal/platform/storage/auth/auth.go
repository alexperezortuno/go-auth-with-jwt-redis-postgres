package auth

import (
	"encoding/json"
	db "github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/data_base"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/data_base/model"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/redis_db"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/tools/common"
	"github.com/golang-jwt/jwt"
	"log"
)

var JWT_TOKEN_SECRET = "ZHTU3oHo6XButkt89ZkJRVUKcWPXDzbLU5UaGA3xYPpY6ASB873GXRJgXQp3pWTATNbNHtufS22xdLYrKf4NqCy5nNaKRryd"
var JWT_REFRESH_TOKEN = "YYELq6utge8Z9C8ynHawaHqcRpV6z33QT2mgfNdhL7porH4VPH3t3ppDSdprpzrMGNSKsmEK4aoFaarNmPByWFytEdLjBsLv"

type CustomClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type TokenResponse struct {
	Status       bool   `json:"status"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

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

func ValidateUser(ar AuthRequest, redisDb *redis_db.Database) (TokenResponse, string) {
	user := model.User{}
	_, err := user.ValidUser(db.Connection, ar.Email, ar.Password)

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return TokenResponse{}, err.Error()
	}

	claimsToken := CustomClaim{
		ar.Email,
		jwt.StandardClaims{
			ExpiresAt: common.NowAdd(1, 0, 0).Unix(),
			Issuer:    "Makobe",
		},
	}

	claimsRefreshToken := CustomClaim{
		ar.Email,
		jwt.StandardClaims{
			Issuer: "Makobe",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsToken)
	signedToken, err := t.SignedString([]byte(JWT_TOKEN_SECRET))
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return TokenResponse{}, err.Error()
	}

	_, _ = redis_db.Set(redisDb, ar.Email, signedToken)
	_, _ = redis_db.Get(redisDb, ar.Email)

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)
	refreshToken, err := rt.SignedString([]byte(JWT_REFRESH_TOKEN))
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return TokenResponse{}, err.Error()
	}

	return TokenResponse{
		true,
		signedToken,
		refreshToken,
	}, ""
}

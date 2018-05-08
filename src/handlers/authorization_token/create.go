package authorization_token

import (
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func New(w http.ResponseWriter, r *http.Request) {
	token, _ := CreateJwt()
	w.Header().Add("Authorization", token)
}

func CreateJwt() (string, error) {
	signingKey := []byte(os.Getenv("JWT_KEY"))
	expiresInSeconds, _ := strconv.Atoi(os.Getenv("JWT_CLAIM_EXPIRES_IN_SECONDS"))
	issuedAt := time.Now()

	claims := &jwt.StandardClaims{
		Id:        uuid.New().String(),
		IssuedAt:  issuedAt.Unix(),
		ExpiresAt: issuedAt.Add(time.Duration(int64(expiresInSeconds)) * time.Second).Unix(),
		Issuer:    os.Getenv("JWT_CLAIMS_ISSUER"),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(signingKey)
}

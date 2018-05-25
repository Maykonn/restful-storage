package authorization_token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"maykonn/restful-storage/src/storage"
	"net/http"
	"os"
	"strconv"
	"time"
)

func New(w http.ResponseWriter, r *http.Request) {
	//c := make(chan string)
	//go CreateJwt(c)

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Server does not support Flusher!",
			http.StatusInternalServerError)
		return
	}

	token, _ := CreateJwt()

	//w.Header().Add("Authorization", <-c)
	w.Header().Add("Authorization", token)
	flusher.Flush()

	storage.Create("")
}

//func CreateJwt(c chan string) {
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
	//tokenString, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(signingKey)
	//c <- tokenString
}

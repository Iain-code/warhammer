package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	pass := []byte(password)
	hashed, err := bcrypt.GenerateFromPassword(pass, 10)
	if err != nil {
		return "", err
	}
	passStr := string(hashed)
	return passStr, nil
}

func CompareHashedPassword(password string, hashed string) error {
	pass := []byte(password)
	hash := []byte(hashed)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		return err
	}
	return nil
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   userID.String(),
		Issuer:    "warhammer",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {

	// Create a new jwt.RegisteredClaims to hold the parsed claims
	claims := &jwt.RegisteredClaims{}

	// tokenString =  The JWT string to parse
	// claims = Where to put the extracted claims
	// func(token *jwt.Token) = This function provides the key to check the signature
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	if !token.Valid {
		return uuid.Nil, errors.New("invalid token")
	}
	sub := claims.Subject // takes the subject field of claims struct (which is the USER.ID)
	if sub == "" {
		return uuid.Nil, err
	}

	notstr, err := uuid.Parse(sub)
	if err != nil {
		return uuid.Nil, err
	}

	return notstr, nil
}

func GetBearerToken(headers http.Header) (string, error) {

	header := headers.Get("Authorization")
	if header == "" {
		return "", errors.New("no authorization header found")
	}
	if !strings.HasPrefix(header, "Bearer ") {
		return "", errors.New("invalid token")
	}

	tkn := strings.TrimSpace(header[7:])
	return tkn, nil
}

func MakeRefreshToken() (string, error) {

	stuff := make([]byte, 32)
	_, err := rand.Read(stuff)
	if err != nil {
		return "", err
	}

	hexStr := hex.EncodeToString(stuff)

	return hexStr, nil

}

// func isTokenExpired(tokenString string) (bool, error) {

// 	parts := strings.Split(tokenString, ".")
// 	if len(parts) != 3 {
// 		return true, fmt.Errorf("invalid token")
// 	}

// 	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
// 	if err != nil {
// 		return true, err
// 	}

// 	claims := &jwt.RegisteredClaims{}
// 	if err := json.Unmarshal(payload, &claims); err != nil {
// 		return true, err
// 	}

// 	currentTime := time.Now().Unix()

// 	return claims.ExpiresAt.Unix() < currentTime, nil
// }

package auth

import (
	"barbershop/creativo/pkg/types"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(u types.User, secret types.JWTSecret) (types.JWTToken, error) {
	stringifiedUserID := strconv.Itoa(int(u.ID))

	claims := types.CustomClaims{
		Roles: u.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    os.Getenv("JWT_ISSUER"), // TODO: Make this configurable
			Subject:   stringifiedUserID,
			ID:        stringifiedUserID,
			Audience:  jwt.ClaimStrings{os.Getenv("JWT_AUDIENCE")}, // TODO: Make this configurable
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return types.JWTToken(tokenString), nil
}

func VerifyJWTToken(tokenString, secretKey string) (*types.CustomClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(tokenString, &types.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*types.CustomClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

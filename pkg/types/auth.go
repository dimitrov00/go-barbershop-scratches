package types

import "github.com/golang-jwt/jwt/v5"

type (
	Password       string
	HashedPassword string
	JWTToken       string
	JWTSecret      string
)

type (
	CustomClaims struct {
		jwt.RegisteredClaims
		Roles []UserRole `json:"roles"`
	}
)

type (
	GenerateJWTTokenFn func(User, JWTSecret) (JWTToken, error)
	VerifyJWTTokenFn   func(tokenString, secretKey string) (*CustomClaims, error)

	HashPasswordFn           func(Password) (HashedPassword, error)
	CompareHashAndPasswordFn func(HashedPassword, Password) error
)

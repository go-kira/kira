package jwt

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-kira/kira"
)

// JWT - Middleware.
type JWT struct{}

// New - return Limitbody instance
func New() *JWT {
	return &JWT{}
}

// CreateToken generate JWT token.
func CreateToken(ctx *kira.Context, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(ctx.Config().GetString("app.key")))
}

// Middleware handler
func (j *JWT) Middleware(ctx *kira.Context, next kira.HandlerFunc) {
	// The token string that will be validated.
	var tokenString string
	// token error
	var err error

	// From where we should grap the token.
	lookup := strings.Split(ctx.Config().GetString("jwt.lookup", "header:Authorization"), ":")

	if lookup[0] == "header" { // From header
		tokenString, err = j.fromHeader(ctx, lookup[1])
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}
	} else if lookup[0] == "cookie" { // From cookie
		tokenString, err = j.fromCookie(ctx, lookup[1])
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}
	}

	// Validate if the request has a valide JWT Token.
	if j.validateToken(ctx, tokenString) {
		next(ctx)
	} else {
		ctx.Status(http.StatusUnauthorized)
		return
	}
}

func (j *JWT) validateToken(ctx *kira.Context, s string) bool {
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(ctx.Config().GetString("app.key")), nil
	})
	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}

	return false
}

// Extract the token from the request header.
func (j *JWT) fromHeader(ctx *kira.Context, key string) (string, error) {
	authHeader := ctx.Request().Header.Get(key)
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("Authorization header format is not valid")
	}

	return authHeaderParts[1], nil
}

// Extract the token from the request cookie.
func (j *JWT) fromCookie(ctx *kira.Context, key string) (string, error) {
	c, err := ctx.Request().Cookie(key)
	if err != nil {
		return "", err
	}

	return c.Value, nil
}

package service

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTManager struct {
	Secret []byte
	Issuer string
	Expiry time.Duration
}

type Claims struct {
	PlayerID string `json:"uid"`
	jwt.RegisteredClaims
}

func (m *JWTManager) Generate(playerId string) (string, error) {
	now := time.Now()
	claims := Claims{
		PlayerID: playerId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.Issuer,
			Subject:   playerId,
			ExpiresAt: jwt.NewNumericDate(now.Add(m.Expiry)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.Secret)
}

// Context key
type ctxKey string

const ctxPlayerID ctxKey = "playerID"

func TryGetPlayerIDFromContext(ctx context.Context) (primitive.ObjectID, bool) {
	v := ctx.Value(ctxPlayerID)
	if s, ok := v.(string); ok {
		id, err := primitive.ObjectIDFromHex(s)
		if err == nil {
			return id, true
		}
	}
	return primitive.NilObjectID, false
}

func (m *JWTManager) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authz := r.Header.Get("Authorization")
		if authz == "" || !strings.HasPrefix(authz, "Bearer ") {
			http.Error(w, "missing bearer token", http.StatusUnauthorized)
			return
		}
		raw := strings.TrimPrefix(authz, "Bearer ")
		claims := &Claims{}
		_, err := jwt.ParseWithClaims(raw, claims, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("invalid signing method")
			}
			return m.Secret, nil
		})
		if err != nil {
			http.Error(w, "invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ctxPlayerID, claims.PlayerID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

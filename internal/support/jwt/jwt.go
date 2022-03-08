package jwt

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"time"
)

type Payload struct {
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.StandardClaims
}

//go:generate mockgen -source=jwt.go -destination=mocks/jwt_mock.go
type Service interface {
	CreateJWT(email, role string) (*string, error)
	VerifyJWT(token string) (*Payload, error)
}

type service struct {
	secretKey   string
	expiry      int
	redisClient *redis.Client
}

func NewJwtService(secretKey string, expiry *int, redisClient *redis.Client) (Service, error) {
	if secretKey == "" {
		return nil, errors.New("invalid jwt secret key")
	}
	if expiry == nil {
		return nil, errors.New("invalid jwt expiry")
	}
	if redisClient == nil {
		return nil, errors.New("invalid redis client")
	}
	return &service{secretKey: secretKey, expiry: *expiry, redisClient: redisClient}, nil
}

func (s *service) CreateJWT(email, role string) (*string, error) {
	// create JWT
	payload := &Payload{
		Email:     email,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Minute * time.Duration(s.expiry)),
	}

	// sign token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := jwtToken.SignedString([]byte(s.secretKey))
	if err != nil {
		return nil, errors.New("token is invalid")
	}

	return &signedToken, nil
}

func (s *service) VerifyJWT(token string) (*Payload, error) {
	// verify JWT
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("token is invalid")
		}
		return []byte(s.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		if _, ok := err.(*jwt.ValidationError); ok {
			return nil, errors.New("token has expired")
		}
		return nil, errors.New("token is invalid")
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errors.New("token is invalid")
	}

	return payload, nil
}
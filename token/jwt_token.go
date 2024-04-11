package token

import (
	"fmt"
	"time"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secret string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invaild key size: must be atleast %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateTokeneateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {

}

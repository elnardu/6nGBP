package model

import (
	"context"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

type TokenData struct {
	ID    bson.ObjectId
	Admin bool
}

func ValidateAndParseToken(ctx context.Context) (*TokenData, error) {
	token, ok := ctx.Value("user").(*jwt.Token)
	if ok != true {
		return nil, errors.New("Unauthorized")
	}

	claims := token.Claims.(jwt.MapClaims)

	dateUnix, ok := claims["exp"].(float64)
	if ok != true {
		return nil, errors.New("Error parsing token")
	}

	date := time.Unix(int64(dateUnix), 0)

	if date.Before(time.Now()) {
		return nil, errors.New("Session has expired")
	}

	var tokenData TokenData

	admin, ok := claims["admin"].(bool)
	if ok != true {
		return nil, errors.New("Error parsing token")
	}

	idString, ok := claims["id"].(string)
	if ok != true {
		return nil, errors.New("Error parsing token")
	}
	tokenData = TokenData{
		Admin: admin,
		ID:    bson.ObjectIdHex(idString),
	}

	return &tokenData, nil
}

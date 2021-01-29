package model

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	jjwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/virgool/blog/conf"
	"github.com/emadghaffari/virgool/blog/database/redis"
)

var (
	// JWT variable instance of intef
	JWT intef = &wt{}
)

// jwt struct
type jwt struct {
	AccessUUID   string `json:"uuid"`
	AtExpires    int64  `json:"exp"`
	RtExpires    int64  `json:"rexp"`
}

type intef interface {
	Get(ctx context.Context, token string, response interface{}) error
	verify(tk string) (string, error)
}
type wt struct{}

// Get, for get values from jwt token
func (j *wt) Get(ctx context.Context, token string, response interface{}) error {
	uuid,err := j.verify(token)
	if err != nil {
		return err
	}

	if err := redis.Database.Get(ctx, uuid, &response); err != nil {
		return err
	}
	return nil
}

// verify the jwt
func (j *wt) verify(tk string) (string, error) {

	token, err := jjwt.Parse(tk, func(token *jjwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jjwt.SigningMethodHMAC); !ok {
			logrus.Warn(fmt.Sprintf("Error in unexpected signing method: %v", token))
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["uuid"])
		}
		return []byte(conf.GlobalConfigs.JWT.Secret), nil
	})
	if err != nil {
		logrus.Warn(fmt.Sprintf("Error in jwt parse: %s", err))
		return "", err
	}
	if _, ok := token.Claims.(jjwt.Claims); !ok && !token.Valid {
		logrus.Warn(fmt.Sprintf("Error in jwt validation: %s", err))

		return "", err
	}

	claims, ok := token.Claims.(jjwt.MapClaims)
	if ok && token.Valid {
		AccessUUID, ok := claims["uuid"].(string)
		if !ok {
			logrus.Warn(fmt.Sprintf("Error in claims uuid from client: %s", err))
			return "", fmt.Errorf("Error in claims uuid from client")
		}

		return AccessUUID, nil
	}
	logrus.Warn(fmt.Sprintf("Error in jwt token verify: %s", err))
	return "", err
}

// Generate hash key
func hasher(lenght int) string {
	letters := []int32("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789-&()_")
	rand.Seed(time.Now().UnixNano())
	b := make([]int32, lenght)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

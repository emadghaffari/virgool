package model

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	jjwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/virgool/auth/conf"
	"github.com/emadghaffari/virgool/auth/database/redis"
)

var (
	// JWT variable instance of intef
	JWT intef = &wt{}
)

// jwt struct
type jwt struct {
	AccessToken  string `json:"at"`
	RefreshToken string `json:"rt"`
	AccessUUID   string `json:"uuid"`
	RefreshUUID  string `json:"rau"`
	AtExpires    int64  `json:"exp"`
	RtExpires    int64  `json:"rexp"`
}

type intef interface {
	Generate(ctx context.Context, model interface{}) (*jwt, error)
	genJWT() (*jwt, error)
	genRefJWT(td *jwt) error
	store(ctx context.Context, model interface{}, td *jwt) error
	Get(ctx context.Context, token string, response interface{}) error
	Verify(tk string) (string, error)
}
type wt struct{}

func (j *wt) Generate(ctx context.Context, model interface{}) (*jwt, error) {

	td, err := j.genJWT()
	if err != nil {
		return nil, err
	}

	if err := j.genRefJWT(td); err != nil {
		return nil, err
	}

	if err := j.store(ctx, model, td); err != nil {
		return nil, err
	}

	return td, nil
}

func (j *wt) genJWT() (*jwt, error) {
	// create new jwt
	td := &jwt{}
	td.AtExpires = time.Now().Add(time.Duration(conf.GlobalConfigs.Service.Redis.UserDuration)).Unix()
	td.RtExpires = time.Now().Add(time.Duration(conf.GlobalConfigs.Service.Redis.UserDuration)).Unix()
	td.AccessUUID = hasher(30)
	td.RefreshUUID = hasher(60)

	// New MapClaims for access token
	atClaims := jjwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["uuid"] = td.AccessUUID
	atClaims["exp"] = td.AtExpires
	at := jjwt.NewWithClaims(jjwt.SigningMethodHS256, atClaims)

	var err error
	td.AccessToken, err = at.SignedString([]byte(conf.GlobalConfigs.JWT.Secret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (j *wt) genRefJWT(td *jwt) error {
	// New MapClaims for refresh access token
	rtClaims := jjwt.MapClaims{}
	rtClaims["uuid"] = td.RefreshUUID
	rtClaims["exp"] = td.RtExpires
	rt := jjwt.NewWithClaims(jjwt.SigningMethodHS256, rtClaims)

	var err error
	td.RefreshToken, err = rt.SignedString([]byte(conf.GlobalConfigs.JWT.RSecret))
	if err != nil {
		return err
	}
	return nil
}

func (j *wt) store(ctx context.Context, model interface{}, td *jwt) error {
	bt, err := json.Marshal(model)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"message": fmt.Sprintf("can not marshal data: %s", model),
			"error":   fmt.Sprintf("Error: %s", err),
		}).Fatal(fmt.Sprintf("can not marshal data: %s", model))
		return fmt.Errorf("can not marshal data: %s", model)
	}
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	now := time.Now()

	// make map for store in redis
	if err := redis.Database.Set(ctx, td.AccessUUID, string(bt), at.Sub(now)); err != nil {
		return err
	}
	return nil
}

func (j *wt) Get(ctx context.Context, token string, response interface{}) error {
	if err := redis.Database.Get(ctx, token, &response); err != nil {
		return err
	}
	return nil
}

func (j *wt) Verify(tk string) (string, error) {
	strArr := strings.Split(tk, " ")
	if len(strArr) != 2 {
		return "", fmt.Errorf("invalid JWT token")
	}
	token, err := jjwt.Parse(strArr[1], func(token *jjwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jjwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.GlobalConfigs.JWT.Secret), nil
	})
	if err != nil {
		return "", err
	}
	if _, ok := token.Claims.(jjwt.Claims); !ok && !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jjwt.MapClaims)
	if ok && token.Valid {
		AccessUUID, ok := claims["uuid"].(string)
		if !ok {
			return "", fmt.Errorf("Error in claims uuid from client")
		}

		return AccessUUID, nil
	}
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

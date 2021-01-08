package model

import (
	"math/rand"
	"time"

	jjwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"

	"github.com/emadghaffari/virgool/auth/conf"
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
	Generate() (*jwt, error)
}
type wt struct{}

func (j *wt) Generate() (*jwt, error) {

	td, err := j.genJWT()
	if err != nil {
		return nil, err
	}

	if err := j.genRefJWT(td); err != nil {
		return nil, err
	}

	return td, nil
}

func (j *wt) genJWT() (*jwt, error) {
	// create new jwt
	td := &jwt{}
	td.AtExpires = time.Now().Add(time.Duration(time.Minute * viper.GetDuration("jwt.expire"))).Unix()
	td.RtExpires = time.Now().Add(time.Duration(time.Minute * viper.GetDuration("jwt.RTexpire"))).Unix()
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

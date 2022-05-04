package model

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	ID     int    `json:"ID"`
	OpenId string `json:"OpenId"`
	jwt.StandardClaims
}

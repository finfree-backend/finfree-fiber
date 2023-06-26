package main

import (
	"github.com/finfree-backend/finfree-fiber/finfiber"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"testing"
	"time"
)

func TestGenerateHS256Token(t *testing.T) {
	mw := finfiber.NewDefaultHS256([]byte("<EXAMPLE-KEY>"))

	exp := time.Now().UTC().Add(time.Hour * 4)
	claims := exampleMapClaims{
		Username: "captain_jack",
		Name:     "Jack",
		Lastname: "Sparrow",
		IsPaid:   true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token, err := mw.GenerateToken(&claims)
	if err != nil {
		log.Println("Unexpected error generating token. Err ->", err)
		t.FailNow()
	}

	log.Println("Token generated successfully. Token ->", token)
}

type exampleMapClaims struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	IsPaid   bool   `json:"is_paid"`
	jwt.RegisteredClaims
}

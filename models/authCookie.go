package models

type Auth struct {
	Token string `cookie:"token"`
}

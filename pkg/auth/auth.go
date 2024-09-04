package auth

import (
	"log"

	"github.com/zalando/go-keyring"
)

type Auth struct {
	service string // name of the application
	user    string // username
}

func New(service, user string) Auth {
	return Auth{
		service: service,
		user:    user,
	}
}

func (a *Auth) SetApiKey(apikey string) error {
	err := keyring.Set(a.service, a.user, apikey)
	if err != nil {
		return err
	}

	log.Println("successfully saved apikey in keyring")
	return nil
}

func (a *Auth) GetAPIKey() (string, error) {
	return keyring.Get(a.service, a.user)
}

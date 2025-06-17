package util

import (
	"errors"
	"os"
)

type Env struct {
	User string
}

func CheckEnv() (*Env, error) {
	user, exists := os.LookupEnv("USER_NAME")
	if !exists {
		return nil, errors.New("USER_NAME environment variable not set")
	}
	env := &Env{
		User: user,
	}
	return env, nil
}

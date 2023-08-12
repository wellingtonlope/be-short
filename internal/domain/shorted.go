package domain

import (
	"errors"
	"math/rand"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRESTUVWXYZ"
)

var (
	ErrUrlIsInvalid = errors.New("url must not be invalid")
)

type Shorted struct {
	Hash string
	Url  string
}

func NewShorted(Url string) (Shorted, error) {
	if Url == "" {
		return Shorted{}, ErrUrlIsInvalid
	}

	Hash := GenerateHash()

	return Shorted{
		Hash: Hash,
		Url:  Url,
	}, nil

}

func GenerateHash() string {
	UrlSize := make([]byte, 6)
	for i := range UrlSize {
		UrlSize[i] = letters[rand.Intn(len(letters))]
	}

	return string(UrlSize)
}

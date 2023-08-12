package domain

import (
	"errors"
	"math/rand"
	"regexp"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRESTUVWXYZ"
)

var (
	ErrUrlIsInvalid = errors.New("url must not be invalid")
)

type Shorted struct {
	Hash string
	URL  string
}

func NewShorted(url string) (Shorted, error) {

	match, _ := regexp.MatchString("^(https?:\\/\\/(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9])(:?\\d*)\\/?([a-z_\\/0-9\\#.]*)\\??([a-z_\\/0-9\\#=&]*)?$", url)

	if !match {
		return Shorted{}, ErrUrlIsInvalid
	}

	hash := GenerateHash()

	return Shorted{
		Hash: hash,
		URL:  url,
	}, nil

}

func GenerateHash() string {
	hashSlice := make([]byte, 6)
	for i := range hashSlice {
		hashSlice[i] = letters[rand.Intn(len(letters))]
	}

	return string(hashSlice)
}

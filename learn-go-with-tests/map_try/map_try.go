package map_try

import (
	"errors"
)

type Dictionary map[string]string

var ErrMsg = "couldn't find the word"

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", errors.New(ErrMsg)
	}

	return value, nil
}

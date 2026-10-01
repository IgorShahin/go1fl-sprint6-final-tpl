package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {
	if data == "" {
		return "", errors.New("data is empty")
	}

	isText := strings.ContainsFunc(data, func(r rune) bool { return r != '.' && r != ' ' && r != '-' })

	if isText {
		return morse.ToMorse(data), nil
	}

	return morse.ToText(data), nil
}

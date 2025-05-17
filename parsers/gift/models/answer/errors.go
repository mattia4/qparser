package giftAnswer

import "errors"

var errMsg = "Formato risposta non valido"

const (
	ErrInvalidFormat = "formato risposta non valido"
	ErrEmptyText     = "il testo della risposta è vuoto"
)

func GetError() error {
	return errors.New(errMsg)
}
